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
type ObjectA struct {
	name string
	value Any
	next *ObjectA
}
type ObjectB map[string]Any

func CreateObject(keys []string, props []Any) Any {
	length := len(keys)
	if length < 1 {
		return &Object0{ }
	} else if length == 1 {
		return &Object1{ k0: keys[0], v0: props[0] }
	} else if length == 2 {
		return &Object2{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
    }
	} else if length == 3 {
		return &Object4{
      k0: keys[0],
      v0: props[0],
      k1: keys[1],
      v1: props[1],
      k2: keys[2],
      v2: props[2],
    }
	} else if length == 4 {
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
	} else if length == 5 {
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
	} else if length == 6 {
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
	} else if length == 7 {
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
	} else if length == 8 {
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
	} else if length == 9 {
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
	} else if length == 10 {
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
	} else if length == 11 {
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
	} else if length == 12 {
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
	} else if length == 13 {
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
	} else if length == 14 {
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
	} else if length == 15 {
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
	} else if length == 16 {
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
	} else if length < 40 {
		obj := &ObjectA{ name: keys[0], value: props[0] }
		for i := 1; i < length; i++ {
			obj.next = &ObjectA{ name: keys[i], value: props[i] }
		}
		return obj
	} else {
    obj := make(ObjectB, length)
    for i := 0; i < length; i++ {
			obj[keys[i]] = props[i]
		}
		return obj
	}
}

func SetValue(obj Any, key string, value Any) {
	if o, ok := obj.(*Object1); ok {
		if o.k0 == key {
			o.v0 = value
		}
	} else if o, ok := obj.(*Object2); ok {
		if o.k0 == key {
      o.v0 = value
    } else if o.k1 == key {
      o.v1 = value
    }
	} else if o, ok := obj.(*Object4); ok {
		if o.k0 == key {
      o.v0 = value
    } else if o.k1 == key {
      o.v1 = value
    } else if o.k2 == key {
      o.v2 = value
    } else if o.k3 == key {
      o.v3 = value
    }
	} else if o, ok := obj.(*Object8); ok {
		if o.k0 == key {
      o.v0 = value
    } else if o.k1 == key {
      o.v1 = value
    } else if o.k2 == key {
      o.v2 = value
    } else if o.k3 == key {
      o.v3 = value
    } else if o.k4 == key {
      o.v4 = value
    } else if o.k5 == key {
      o.v5 = value
    } else if o.k6 == key {
      o.v6 = value
    } else if o.k7 == key {
      o.v7 = value
    }
	} else if o, ok := obj.(*Object16); ok {
		if o.k0 == key {
      o.v0 = value
    } else if o.k1 == key {
      o.v1 = value
    } else if o.k2 == key {
      o.v2 = value
    } else if o.k3 == key {
      o.v3 = value
    } else if o.k4 == key {
      o.v4 = value
    } else if o.k5 == key {
      o.v5 = value
    } else if o.k6 == key {
      o.v6 = value
    } else if o.k7 == key {
      o.v7 = value
    } else if o.k8 == key {
      o.v8 = value
    } else if o.k9 == key {
      o.v9 = value
    } else if o.k10 == key {
      o.v10 = value
    } else if o.k11 == key {
      o.v11 = value
    } else if o.k12 == key {
      o.v12 = value
    } else if o.k13 == key {
      o.v13 = value
    } else if o.k14 == key {
      o.v14 = value
    } else if o.k15 == key {
      o.v15 = value
    }
	} else if o, ok := obj.(*ObjectA); ok {
		var oo Any = o
		for oo != nil {
			o = oo.(*ObjectA)
			if o.name == key {
				o.value = value
				break
			} else {
				oo = o.next
			}
		}
	} else if o, ok := obj.(*ObjectB); ok {
		(*o)[key] = value
	}
}

func GetValue(obj Any, key string) Any {
	if o, ok := obj.(*Object1); ok {
		if o.k0 == key {
			return o.v0
		}
	} else if o, ok := obj.(*Object2); ok {
		if o.k0 == key {
      return o.v0
    } else if o.k1 == key {
      return o.v1
    }
	} else if o, ok := obj.(*Object4); ok {
		if o.k0 == key {
      return o.v0
    } else if o.k1 == key {
      return o.v1
    } else if o.k2 == key {
      return o.v2
    } else if o.k3 == key {
      return o.v3
    }
	} else if o, ok := obj.(*Object8); ok {
		if o.k0 == key {
      return o.v0
    } else if o.k1 == key {
      return o.v1
    } else if o.k2 == key {
      return o.v2
    } else if o.k3 == key {
      return o.v3
    } else if o.k4 == key {
      return o.v4
    } else if o.k5 == key {
      return o.v5
    } else if o.k6 == key {
      return o.v6
    } else if o.k7 == key {
      return o.v7
    }
	} else if o, ok := obj.(*Object16); ok {
		if o.k0 == key {
      return o.v0
    } else if o.k1 == key {
      return o.v1
    } else if o.k2 == key {
      return o.v2
    } else if o.k3 == key {
      return o.v3
    } else if o.k4 == key {
      return o.v4
    } else if o.k5 == key {
      return o.v5
    } else if o.k6 == key {
      return o.v6
    } else if o.k7 == key {
      return o.v7
    } else if o.k8 == key {
      return o.v8
    } else if o.k9 == key {
      return o.v9
    } else if o.k10 == key {
      return o.v10
    } else if o.k11 == key {
      return o.v11
    } else if o.k12 == key {
      return o.v12
    } else if o.k13 == key {
      return o.v13
    } else if o.k14 == key {
      return o.v14
    } else if o.k15 == key {
      return o.v15
    }
	} else if o, ok := obj.(*ObjectA); ok {
		var oo Any = o
		for oo != nil {
			o = oo.(*ObjectA)
			if o.name == key {
				return o.value
			} else {
				oo = o.next
			}
		}
	} else if o, ok := obj.(*ObjectB); ok {
		return (*o)[key]
  }
  return nil
}
