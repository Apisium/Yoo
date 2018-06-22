package yoo

import (
  "fmt"
  "errors"
	"io/ioutil"
)

var modules = *GetBindings()
var global = GetGlobal()

func Execute(object Any, variables *Variables, variables2 *Variables) (Any, error) {
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
    callee, err := Execute(o.callee, variables, variables2)
    if err != nil { return nil, err }
    argsA := *o.args
    length := len(argsA)
    args := make([]Any, length, length)
    for k, v := range argsA {
      arg, err := Execute(v, variables, variables2)
      if err != nil { return nil, err }
      args[k] = arg
    }
    if fn, ok := callee.(*TSFunction); ok {
      a := (*fn)(&args)
      return a, nil
    }
    return nil, errors.New(fmt.Sprintf("%+v is not a function.", callee))
  case *Member:
    left, err := Execute(o.left, variables, variables2)
    if err != nil { return nil, err }
    rightA, ok := o.right.(*Identifier)
    var right string
    if ok {
      right = rightA.text
    } else {
      ret, err := Execute(o.right, variables, variables2)
      if err != nil { return nil, err }
      right = ret.(string)
    }
    return GetValue(left, right), nil
  case *Variable:
    vars := *variables
    for _, elem := range (*o) {
      v, err := Execute(elem.value, variables, variables2)
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
        Execute(v, vars, variables)
      }
      return nil
    }
    return &fn, nil
  case *TSClass:
    cls := CreateObjectByLength(o.length)
    (*o.constructor)(&[]Any { cls })
    return cls, nil
  }
  return nil, nil
}

func ExecuteFile(file string) error {
	f, err := ioutil.ReadFile("a.yoo")
  if err != nil { return err }
  return ExecuteBytes(f)
}

func ExecuteBytes(buff []byte) error {
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
    _, err = Execute(obj, vars, nil)
    if err != nil { return err }
  }

  return nil
}
