// I don't know how to rewrite this. But, this is the best resolution after my test.

package yoo

type Object0 struct { }
type Object1 struct {
  k0 string
  v0 Any
}
type Object2 struct {
  k0 string
  v0 Any
  k1 string
  v1 Any
}
type Object4 struct {
  k0 string
  v0 Any
  k1 string
  v1 Any
  k2 string
  v2 Any
  k3 string
  v3 Any
}
type Object8 struct {
  k0 string
  v0 Any
  k1 string
  v1 Any
  k2 string
  v2 Any
  k3 string
  v3 Any
  k4 string
  v4 Any
  k5 string
  v5 Any
  k6 string
  v6 Any
  k7 string
  v7 Any
}
type Object16 struct {
  k0 string
  v0 Any
  k1 string
  v1 Any
  k2 string
  v2 Any
  k3 string
  v3 Any
  k4 string
  v4 Any
  k5 string
  v5 Any
  k6 string
  v6 Any
  k7 string
  v7 Any
  k8 string
  v8 Any
  k9 string
  v9 Any
  k10 string
  v10 Any
  k11 string
  v11 Any
  k12 string
  v12 Any
  k13 string
  v13 Any
  k14 string
  v14 Any
  k15 string
  v15 Any
}
type ObjectA map[string]Any
type ProxyFunction func (object Any, t int, key string, value Any) Any
type Proxy struct {
  object Any
  function *ProxyFunction
}

func CreateObjectByLength(length int) Any {
  switch {
  case length < 1:
    return &Object0{}
  case length == 1:
    return &Object1{}
  case length == 2:
    return &Object2{}
  case length < 5:
    return &Object4{}
  case length < 9:
    return &Object8{}
  case length < 17:
    return &Object16{}
  default:
    obj := make(ObjectA, length)
    return &obj
  }
}

func CreateObject(keysPtr *[]string, propsPtr *[]Any) Any {
  keys := *keysPtr
  props := *propsPtr
  length := len(keys)

  if length < 1 { return &Object0{ } }
  switch length {
  case 1:
    return &Object1{ k0: keys[0], v0: props[0] }
  case 2:
    return &Object2{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
    }
  case 3:
    return &Object4{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
    }
  case 4:
    return &Object4{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
    }
  case 5:
    return &Object8{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
    }
  case 6:
    return &Object8{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
    }
  case 7:
    return &Object8{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
    }
  case 8:
    return &Object8{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
      k7: keys[7],
      v7: props[7],
    }
  case 9:
    return &Object16{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
      k7: keys[7],
      v7: props[7],
      k8: keys[8],
      v8: props[8],
    }
  case 10:
    return &Object16{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
      k7: keys[7],
      v7: props[7],
      k8: keys[8],
      v8: props[8],
      k9: keys[9],
      v9: props[9],
    }
  case 11:
    return &Object16{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
      k7: keys[7],
      v7: props[7],
      k8: keys[8],
      v8: props[8],
      k9: keys[9],
      v9: props[9],
      k10: keys[10],
      v10: props[10],
    }
  case 12:
    return &Object16{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
      k7: keys[7],
      v7: props[7],
      k8: keys[8],
      v8: props[8],
      k9: keys[9],
      v9: props[9],
      k10: keys[10],
      v10: props[10],
      k11: keys[11],
      v11: props[11],
    }
  case 13:
    return &Object16{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
      k7: keys[7],
      v7: props[7],
      k8: keys[8],
      v8: props[8],
      k9: keys[9],
      v9: props[9],
      k10: keys[10],
      v10: props[10],
      k11: keys[11],
      v11: props[11],
      k12: keys[12],
      v12: props[12],
    }
  case 14:
    return &Object16{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
      k7: keys[7],
      v7: props[7],
      k8: keys[8],
      v8: props[8],
      k9: keys[9],
      v9: props[9],
      k10: keys[10],
      v10: props[10],
      k11: keys[11],
      v11: props[11],
      k12: keys[12],
      v12: props[12],
      k13: keys[13],
      v13: props[13],
    }
  case 15:
    return &Object16{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
      k7: keys[7],
      v7: props[7],
      k8: keys[8],
      v8: props[8],
      k9: keys[9],
      v9: props[9],
      k10: keys[10],
      v10: props[10],
      k11: keys[11],
      v11: props[11],
      k12: keys[12],
      v12: props[12],
      k13: keys[13],
      v13: props[13],
      k14: keys[14],
      v14: props[14],
    }
  case 16:
    return &Object16{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
      k3: keys[3],
      v3: props[3],
      k4: keys[4],
      v4: props[4],
      k5: keys[5],
      v5: props[5],
      k6: keys[6],
      v6: props[6],
      k7: keys[7],
      v7: props[7],
      k8: keys[8],
      v8: props[8],
      k9: keys[9],
      v9: props[9],
      k10: keys[10],
      v10: props[10],
      k11: keys[11],
      v11: props[11],
      k12: keys[12],
      v12: props[12],
      k13: keys[13],
      v13: props[13],
      k14: keys[14],
      v14: props[14],
      k15: keys[15],
      v15: props[15],
    }
  default:
    obj := make(ObjectA, length)
    for i := 0; i < length; i++ {
      obj[keys[i]] = props[i]
    }
    return &obj
  }
}

func SetValue(obj Any, key string, value Any) {
  switch o := obj.(type) {
  case *Object1:
    if o.k0 == key {
      o.v0 = value
    }
  case *Object2:
    switch key {
    case o.k0:
      o.v0 = value
    case o.k1:
      o.v1 = value
    }
  case *Object4:
    switch key {
    case o.k0:
      o.v0 = value
    case o.k1:
      o.v1 = value
    case o.k2:
      o.v2 = value
    case o.k3:
      o.v3 = value
    }
  case *Object8:
    switch key {
    case o.k0:
      o.v0 = value
    case o.k1:
      o.v1 = value
    case o.k2:
      o.v2 = value
    case o.k3:
      o.v3 = value
    case o.k4:
      o.v4 = value
    case o.k5:
      o.v5 = value
    case o.k6:
      o.v6 = value
    case o.k7:
      o.v7 = value
    }
  case *Object16:
    switch key {
    case o.k0:
      o.v0 = value
    case o.k1:
      o.v1 = value
    case o.k2:
      o.v2 = value
    case o.k3:
      o.v3 = value
    case o.k4:
      o.v4 = value
    case o.k5:
      o.v5 = value
    case o.k6:
      o.v6 = value
    case o.k7:
      o.v7 = value
    case o.k8:
      o.v8 = value
    case o.k9:
      o.v9 = value
    case o.k10:
      o.v10 = value
    case o.k11:
      o.v11 = value
    case o.k12:
      o.v12 = value
    case o.k13:
      o.v13 = value
    case o.k14:
      o.v14 = value
    case o.k15:
      o.v15 = value
    }
  case *ObjectA:
    (*o)[key] = value
  case *Proxy:
    (*o.function)(o.object, PROXY_SET, key, value)
  }
}

func GetValue(obj Any, key string) Any {
  switch o := obj.(type) {
  case *Object1:
    if o.k0 == key {
      return o.v0
    }
  case *Object2:
    switch key {
    case o.k0:
      return o.v0
    case o.k1:
      return o.v1
    }
  case *Object4:
    switch key {
    case o.k0:
      return o.v0
    case o.k1:
      return o.v1
    case o.k2:
      return o.v2
    case o.k3:
      return o.v3
    }
  case *Object8:
    switch key {
    case o.k0:
      return o.v0
    case o.k1:
      return o.v1
    case o.k2:
      return o.v2
    case o.k3:
      return o.v3
    case o.k4:
      return o.v4
    case o.k5:
      return o.v5
    case o.k6:
      return o.v6
    case o.k7:
      return o.v7
    }
  case *Object16:
    switch key {
    case o.k0:
      return o.v0
    case o.k1:
      return o.v1
    case o.k2:
      return o.v2
    case o.k3:
      return o.v3
    case o.k4:
      return o.v4
    case o.k5:
      return o.v5
    case o.k6:
      return o.v6
    case o.k7:
      return o.v7
    case o.k8:
      return o.v8
    case o.k9:
      return o.v9
    case o.k10:
      return o.v10
    case o.k11:
      return o.v11
    case o.k12:
      return o.v12
    case o.k13:
      return o.v13
    case o.k14:
      return o.v14
    case o.k15:
      return o.v15
    }
  case *ObjectA:
    return (*o)[key]
  case *Proxy:
    return (*o.function)(o.object, PROXY_GET, key, nil)
  }
  return nil
}

func IsObject(obj Any) bool {
  switch obj.(type) {
  case *Object0:
    return true
  case *Object1:
    return true
  case *Object2:
    return true
  case *Object4:
    return true
  case *Object8:
    return true
  case *Object16:
    return true
  case *ObjectA:
    return true
  case *Proxy:
    return true
  default:
    return false
  }
}
