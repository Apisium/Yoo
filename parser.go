package yoo

import (
  "fmt"
  "time"
  "errors"
)

var pool []string

type Variables map[string]Any
type TSFunction func(args *[]Any) Any
var global = make(Variables)

func init() {
  var log TSFunction = func (args *[]Any) Any {
    for _, v := range *args {
      if o, ok := v.(*Any); ok {
        fmt.Print(*o, " ")
      } else {
        fmt.Print(v, " ")
      }
    }
    fmt.Println()
    return nil
  }
  global["console"] = CreateObject(
    []string{ "log" },
    []Any{ &log },
  )
}

func findObject(object Any, variables *Variables) Any {
  if obj1, ok1 := object.(*Any); ok1 {
    obj := *obj1
    if o, ok := obj.(*Identifier); ok {
      return (*variables)[o.text]
    } else if o, ok := obj.(*String); ok {
      return pool[o.id]
    }
  } else if fn, ok1 := object.(*TSFunction); ok1 {
    return fn
  }
  return nil
}

func expression(ba *ByteArray, variables *Variables) (Any, error) {
  ret, err := ba.ReadInt8()
  if err != nil { return nil, err }

  if ret == CONSTANT_IDENTIFIER { return NewIdentifier(ba) }
  if ret == CONSTANT_STRING { return NewString(ba) }
  if ret == CONSTANT_CALL {
    fn, err := NewCall(ba, variables)
    if err != nil { return nil, err }
    callee := findObject(fn.callee, variables)

    argsA := *fn.args
    length := len(argsA)
    args := make([]Any, length, length)
    for k, v := range argsA {
      args[k] = findObject(v, variables)
    }
    return (*callee.(*TSFunction))(&args), nil
  }
  if ret == CONSTANT_MEMBER {
    member, err := NewMember(ba, variables)
    if err != nil { return nil, err }
    left := findObject(member.left, variables)
    rightA, ok := member.right.(*Identifier)
    var right string
    if ok {
      right = rightA.text
    } else {
      right = findObject(member.right, variables).(string)
    }
    return GetValue(left, right), nil
  }
  if ret == CONSTANT_VARIABLE {
    list, err := NewVariable(ba, variables)
    if err != nil { return nil, err }
    vars := *variables
    for _, elem := range (*list) {
      vars[elem.name.text] = elem.value
    }
    return nil, nil
  }
  if ret == CONSTANT_NULL {
    return nil, nil
  }

  return nil, errors.New("No such token: " + string(ret))
}

func Parse(buff []byte) (err error) {
  defer trace()()
  ba := CreateByteArray(buff)
  head, err := ba.ReadString(3)
  if err != nil || head != "YOO" {
    return errors.New("Header error")
  }

  version, _ := ba.ReadInt8()
  if version != VERSION {
    return errors.New("Version error")
  }
  // fmt.Println("Version:", version)

  ret, err := ba.ReadInt8()
  if err != nil || ret != CONSTANT_STRINGS_START {
    return errors.New("Strings pool error")
  }

  length, err := ba.ReadInt16()
  if err != nil {
    return errors.New("Strings pool error")
  }
  // fmt.Println("String pool length:", length)

  pool = make([]string, length, length)
  for i := int16(0); i < length; i++ {
    len, err := ba.ReadInt64()
    if err != nil {
      return err
    }
    str, err := ba.ReadString(int(len))
    if err != nil {
      return err
    }
    pool[i] = str
  }
  code, err := ba.ReadInt8()
  if err != nil || code != CONSTANT_STRINGS_END {
    return
  }

  vars := copyMap(&global)
  for ba.Available() > 0 {
    expression(ba, vars)
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
