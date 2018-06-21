package yoo

type Block []Any

type Identifier struct {
  text string
}
func NewIdentifier(ba *ByteArray, pool *[]string) (*Identifier, error) {
  id, err := ba.ReadInt16()
  if err != nil { return nil, err }
  return &Identifier{ text: (*pool)[int(id)] }, nil
}

type Call struct {
  callee Any
  args *[]Any
}
func NewCall(ba *ByteArray, pool *[]string) (*Call, error) {
  callee, err := expression(ba, pool)
  if err != nil { return nil, err }
  length, err := ba.ReadInt16()
  if err != nil { return nil, err }

  args := make([]Any, length, length)
  for i := int16(0); i < length; i++ {
    arg, err := expression(ba, pool)
    if err != nil { return nil, err }
    args[i] = &arg
  }
  return &Call{ callee: callee, args: &args }, nil
}

type Member struct {
  left Any
  right Any
}
func NewMember(ba *ByteArray, pool *[]string) (*Member, error) {
  left, err := expression(ba, pool)
  if err != nil { return nil, err }
  right, err := expression(ba, pool)
  if err != nil { return nil, err }
  return &Member{ left: left, right: right }, nil
}

type VariableElement struct {
  name *Identifier
  value Any
}
type Variable []*VariableElement
func NewVariable(ba *ByteArray, pool *[]string) (*Variable, error) {
  length, err := ba.ReadInt16()
  if err != nil { return nil, err }
  vars := make(Variable, length, length)
  for i := int16(0); i < length; i++ {
    left, err := expression(ba, pool)
    if err != nil { return nil, err }
    right, err := expression(ba, pool)
    if err != nil { return nil, err }
    vars[i] = &VariableElement{ name: left.(*Identifier), value: right }
  }
  return &vars, nil
}

type Imported struct {
  name *Identifier
  prop *Identifier
}
type Import struct {
  path string
  importeds *[]*Imported
}
func NewImport(ba *ByteArray, pool *[]string) (*Import, error) {
  path, err := expression(ba, pool)
  if err != nil { return nil, err }
  length, err := ba.ReadInt16()
  if err != nil { return nil, err }

  imports := make([]*Imported, length, length)
  for i := int16(0); i < length; i++ {
    name, err := expression(ba, pool)
    if err != nil { return nil, err }
    prop, err := expression(ba, pool)
    if err != nil { return nil, err }
    imports[i] = &Imported{ name: name.(*Identifier), prop: prop.(*Identifier) }
  }
  return &Import{ path: path.(string), importeds: &imports }, nil
}

type ArrowFunction struct {
  async bool
  args *Variable
  body *Block
}
func NewArrowFunction(ba *ByteArray, pool *[]string) (*ArrowFunction, error) {
  async, err := expression(ba, pool)
  if err != nil { return nil, err }
  args, err := NewVariable(ba, pool)
  if err != nil { return nil, err }
  length, err := ba.ReadInt16()
  if err != nil { return nil, err }

  block := make(Block, length, length)
  for i := int16(0); i < length; i++ {
    expr, err := expression(ba, pool)
    if err != nil { return nil, err }
    block[i] = expr
  }
  return &ArrowFunction{ async: async.(bool), args: args, body: &block }, nil
}
