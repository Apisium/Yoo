package yoo

import (
  "fmt"
  "time"
  "errors"
  "strconv"
)

type Variables map[string]Any
type TSFunction func(args *[]Any) Any
var modulesPtr = GetBindings()
var global = GetGlobal(modulesPtr)
var modules = *modulesPtr

func init() {

}

func expression(ba *ByteArray, pool *[]string) (Any, error) {
  ret, err := ba.ReadInt8()
  if err != nil { return nil, err }

  switch ret {
  case CONSTANT_TRUE:
    return true, nil
  case CONSTANT_FALSE:
    return false, nil
  case CONSTANT_STRING:
    id, err := ba.ReadInt16()
    if err != nil { return nil, err }
    return (*pool)[int(id)], nil
  case CONSTANT_NUMBER:
    length, err := ba.ReadInt16()
    if err != nil { return nil, err }
    text, err := ba.ReadString(int(length))
    if err != nil { return nil, err }
    return strconv.ParseFloat(text, 64)
  case CONSTANT_IDENTIFIER:
    return NewIdentifier(ba, pool)
  case CONSTANT_CALL:
    return NewCall(ba, pool)
  case CONSTANT_MEMBER:
    return NewMember(ba, pool)
  case CONSTANT_VARIABLE:
    return NewVariable(ba, pool)
  case CONSTANT_IMPORT:
    return NewImport(ba, pool)
  case CONSTANT_NULL:
    return nil, nil
  case CONSTANT_ARROW_FUNCTION:
    return NewArrowFunction(ba, pool)
  default:
    return nil, errors.New("No such token: 0x" + strconv.FormatInt(int64(ret), 16))
  }
}

func execute(object Any, variables *Variables, variables2 *Variables) (Any, error) {
  obj := object
  if obj1, ok1 := object.(*Any); ok1 {
    obj = *obj1
  }
  switch o := obj.(type) {
  case string:
    return o, nil
  case float64:
    return o, nil
  case bool:
    return o, nil
  case *Identifier:
    value := (*variables)[o.text]
    if value == nil && variables2 != nil {
      value = (*variables2)[o.text]
    }
    return value, nil
  case *TSFunction:
    return o, nil
  case *Proxy:
    return o, nil
  case *Call:
    callee, err := execute(o.callee, variables, variables2)
    if err != nil { return nil, err }
    argsA := *o.args
    length := len(argsA)
    args := make([]Any, length, length)
    for k, v := range argsA {
      arg, err := execute(v, variables, variables2)
      if err != nil { return nil, err }
      args[k] = arg
    }
    if fn, ok := callee.(*TSFunction); ok {
      a := (*fn)(&args)
      return a, nil
    }
    return nil, errors.New(fmt.Sprintf("%+v is not a function.", callee))
  case *Member:
    left, err := execute(o.left, variables, variables2)
    if err != nil { return nil, err }
    rightA, ok := o.right.(*Identifier)
    var right string
    if ok {
      right = rightA.text
    } else {
      ret, err := execute(o.right, variables, variables2)
      if err != nil { return nil, err }
      right = ret.(string)
    }
    return GetValue(left, right), nil
  case *Variable:
    vars := *variables
    for _, elem := range (*o) {
      v, err := execute(elem.value, variables, variables2)
      if err != nil { return nil, err }
      vars[elem.name.text] = v
    }
  case *Import:
    name := o.path
    module := modules[name]
    if module == nil { return nil, errors.New("No such module: " + name) }
    vars := *variables
    for _, elem := range (*o.importeds) {
      vars[elem.name.text] = GetValue(module, elem.prop.text)
    }
  case *ArrowFunction:
    pars := *o.args
    var fn TSFunction = func (args *[]Any) Any {
      vars := copyMap(variables)
      varsV := *vars
      argsV := *args
      length := len(argsV) - 1
      for i, v := range pars {
        var value Any
        if i <= length {
          value = argsV[i]
        }
        if value == nil {
          value = v.value
        }
        varsV[v.name.text] = value
      }
      for _, v := range (*o.body) {
        execute(v, vars, variables)
      }
      return nil
    }
    return &fn, nil
  }
  return nil, nil
}

func ImportFile(buff []byte) (err error) {
  ba := CreateByteArray(buff)
  head, err := ba.ReadString(3)
  if err != nil || head != "YOO" {
    return errors.New("Header error")
  }

  version, _ := ba.ReadInt8()
  if version != VERSION {
    return errors.New("Version error")
  }

  length, err := ba.ReadInt16()
  if err != nil {
    return errors.New("Strings pool error")
  }

  pool := make([]string, length, length)
  for i := int16(0); i < length; i++ {
    len, err := ba.ReadInt16()
    if err != nil {
      return err
    }
    str, err := ba.ReadString(int(len))
    if err != nil {
      return err
    }
    pool[i] = str
  }

  length, err = ba.ReadInt16()
  if err != nil {
    return errors.New("Expression error")
  }

  poolPtr := &pool
  vars := copyMap(global)
  length--
  for length > 0 {
    length--
    obj, err := expression(ba, poolPtr)
    if err != nil { return err }
    _, err = execute(obj, vars, nil)
    if err != nil { return err }
  }

  return nil
}

func trace() func() {
  start := time.Now()
  return func() {
    fmt.Printf("Cost: %s\n", time.Since(start))
  }
}

func copyMap(obj *Variables) *Variables {
  oldMap := *obj
  newMap := make(Variables, len(oldMap))
  for key, value := range oldMap {
    newMap[key] = value
  }
  return &newMap
}
