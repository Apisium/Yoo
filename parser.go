package yoo

import (
  "fmt"
  "time"
  "errors"
  "strconv"
  "net/http"
)

type Variables map[string]Any
type TSFunction func(args *[]Any) Any
var global = make(Variables)
var modules = make(Variables)

func init() {
  var log TSFunction = func (args *[]Any) Any {
    for _, v := range *args {
      fmt.Print(v, " ")
    }
    fmt.Println()
    return nil
  }
  global["console"] = CreateObject(
    []string{ "log" },
    []Any{ &log },
  )

  var Server TSFunction = func (args *[]Any) Any {
    handler := *(*args)[0].(*TSFunction)
    var h http.HandlerFunc = func (w http.ResponseWriter, r *http.Request) {
      defer trace()()
      handler(nil)
      fmt.Fprint(w, "Hello World!")
    }
    server := http.Server{
      Addr: "127.0.0.1:1234",
      Handler: h,
    }
    server.ListenAndServe()
    return nil
  }
  modules["http"] = CreateObject(
    []string{ "Server" },
    []Any{ &Server },
  )
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

func execute(object Any, variables *Variables) (Any, error) {
  obj := object
  if obj1, ok1 := object.(*Any); ok1 {
    obj = *obj1
  }
  switch o := obj.(type) {
  case string:
    return o, nil
  case bool:
    return o, nil
  case *Identifier:
    return (*variables)[o.text], nil
  case *TSFunction:
    return o, nil
  case *Call:
    callee, err := execute(o.callee, variables)
    if err != nil { return nil, err }
    argsA := *o.args
    length := len(argsA)
    args := make([]Any, length, length)
    for k, v := range argsA {
      arg, err := execute(v, variables)
      if err != nil { return nil, err }
      args[k] = arg
    }
    return (*callee.(*TSFunction))(&args), nil
  case *Member:
    left, err := execute(o.left, variables)
    if err != nil { return nil, err }
    rightA, ok := o.right.(*Identifier)
    var right string
    if ok {
      right = rightA.text
    } else {
      ret, err := execute(o.right, variables)
      if err != nil { return nil, err }
      right = ret.(string)
    }
    return GetValue(left, right), nil
  case *Variable:
    vars := *variables
    for _, elem := range (*o) {
      v, err := execute(elem.value, variables)
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
    vars := copyMap(variables)
    var fn TSFunction = func (args *[]Any) Any {
      for _, v := range (*o.body) {
        execute(v, vars)
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
  vars := copyMap(&global)
  length--
  for length > 0 {
    length--
    obj, err := expression(ba, poolPtr)
    if err != nil { return err }
    _, err = execute(obj, vars)
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
