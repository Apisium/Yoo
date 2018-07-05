package yoo

import (
  "errors"
  "strconv"
	"io/ioutil"
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

func ParseFile(file string) (*Block, error) {
	f, err := ioutil.ReadFile(file)
  if err != nil { return nil, err }
  return ParseBytes(f)
}

func ParseBytes(buff []byte) (ret Any, err error) {
  ba := CreateByteArray(buff)
  head, err := ba.ReadString(3)
  if err != nil || head != "YOO" {
    return errors.New("Header error")
  }

  version, _ := ba.ReadInt8()
  if version != VERSION {
    return errors.New("Version error")
  }

  length, _ := ba.ReadLength()
  if err != nil {
    return errors.New("Strings pool error")
  }

  pool := make([]string, length, length)
  for i := 0; i < length; i++ {
    len, err = ba.ReadLength()
    if err != nil {
      return
    }
    str, err = ba.ReadString(len)
    if err != nil { return }
    pool[i] = str
  }

  for {
    t, err = ba.ReadInt8()
    if err != nil { return }

    switch t {
    case CONSTANT_BLOCK:
      length, err = ba.ReadLength()
      if err != nil { return }
      block := make(Block, length, length)
      for i := int16(0); i < length; i++ {
        expr, err := expression(ba, pool)
        if err != nil { return nil, err }
        block[i] = expr
      }
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
    length, err := ba.ReadLength()
    if err != nil {
      return nil, errors.New("Strings pool error")
    }
  }
  return NewBlock(ba, &pool)
}
