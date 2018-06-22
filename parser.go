package yoo

import (
  "fmt"
  "time"
  "errors"
  "strconv"
)

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
  case CONSTANT_NEW:
    return NewMakeClass(ba, pool)
  default:
    return nil, errors.New("No such token: 0x" + strconv.FormatInt(int64(ret), 16))
  }
}

func T() Any {
  defer trace()()
  var e Any
  for i := 0; i < 100000000; i++ {
    // CreateObjectByLength(2)
    e = make(Variables, 2)
  }
  return e
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
