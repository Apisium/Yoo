package yoo

import (
	"fmt"
	"time"
	"errors"
	// "strings"
)

const (
	VERSION = 0x01
	CONSTANT_CUT = 0x00
	CONSTANT_STRINGS_START = 0x10
  CONSTANT_STRINGS_END = 0x11
  CONSTANT_IDENTIFIER = 0x12
  CONSTANT_CALL = 0x13
	CONSTANT_STRING = 0x14
	PRIMITIVE_VALUE = "[[PRIMITIVE VALUE]]"
)

var pool []string

type TSObject *map[string]Any
var global = make(map[string]TSObject)
type TSFunction func(args []TSObject)

func findObject(obj Any) TSObject {
	switch obj.(type) {
	case *Identifier:
		return global[obj.(*Identifier).text]
	case *String:
		var str = make(map[string]Any)
		str[PRIMITIVE_VALUE] = pool[obj.(*String).id]
		return &str
	}
	return nil
}
func init() {
	var log = make(map[string]Any)
	var logFn TSFunction = func (args []TSObject) {
		fmt.Println((*args[0])[PRIMITIVE_VALUE].(string))
		// length := len(args)
		// data := make([]string, length, length)
		// for i, v := range args { data[i] = (*v)[PRIMITIVE_VALUE].(string) }
		// fmt.Println(strings.Join(data, " "))
	}
	log[PRIMITIVE_VALUE] = logFn
	global["log"] = &log
}

func expr(ba *ByteArray) (Any, error) {
	ret, err := ba.ReadInt8()
	if err != nil { return nil, err }

	if ret == CONSTANT_IDENTIFIER { return NewIdentifier(ba) }
	if ret == CONSTANT_STRING { return NewString(ba) }
	if ret == CONSTANT_CALL {
		fn, err := NewFunction(ba)
		if err != nil { return nil, err }
		callee := *findObject(fn.callee)

		length := len(fn.args)
		args := make([]TSObject, length, length)
		for k, v := range fn.args { args[k] = findObject(v) }
		callee[PRIMITIVE_VALUE].(TSFunction)(args)
	}
	return nil, nil
}

func Parse(buff []byte) (err error) {
	defer trace()()
	ba := CreateByteArray(buff)
	head, err := ba.ReadString(5)
	if err != nil || head != "YOOOO" {
		return errors.New("Header error")
	}

	_, err = ba.ReadInt8()
	if err != nil {
		return
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

	for ba.Available() > 0 {
		expr(ba)
	}

	return nil
}

func trace() func() {
  start := time.Now()
  return func() {
    fmt.Printf("Cost: %s", time.Since(start))
  }
}
