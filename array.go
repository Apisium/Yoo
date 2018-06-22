package yoo

type Array []Any

func getUint8Array() *TSClass {
 var c TSFunction = func (args *[]Any) Any {
    return nil
  }
  return &TSClass{
    constructor: &c,
    length: 0,
  }
}

func getArrays() *Variables {
  return &Variables{
    "Uint8Array": getUint8Array(),
  }
}

// type Uint8Array []uint8

// func NewUint8Array(len Any)
