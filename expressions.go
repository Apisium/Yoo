package yoo

type Identifier struct {
	text string
}
func NewIdentifier(ba *ByteArray) (*Identifier, error) {
	ret, err := ba.ReadInt64()
	if err != nil { return nil, err }
	text, err := ba.ReadString(int(ret))
	if err != nil { return nil, err }
	return &Identifier{ text: text }, nil
}

type String struct {
	id int
}
func NewString(ba *ByteArray) (*String, error) {
	id, err := ba.ReadInt16()
	if err != nil { return nil, err }
	return &String{ id: int(id) }, nil
}

type Call struct {
	callee Any
	args *[]Any
}
func NewCall(ba *ByteArray, variables *Variables) (*Call, error) {
	callee, err := expression(ba, variables)
	if err != nil { return nil, err }
	length, err := ba.ReadInt16()
	if err != nil { return nil, err }

	pool := make([]Any, length, length)
	for i := int16(0); i < length; i++ {
		arg, err := expression(ba, variables)
		if err != nil { return nil, err }
		pool[i] = &arg
	}
	return &Call{ callee: callee, args: &pool }, nil
}

type Member struct {
	left Any
	right Any
}
func NewMember(ba *ByteArray, variables *Variables) (*Member, error) {
	left, err := expression(ba, variables)
	if err != nil { return nil, err }
	right, err := expression(ba, variables)
	if err != nil { return nil, err }
	return &Member{ left: &left, right: &right }, nil
}

type VariableElement struct {
	name *Identifier
	value Any
}
type Variable []*VariableElement
func NewVariable(ba *ByteArray, variables *Variables) (*Variable, error) {
	length, err := ba.ReadInt16()
	if err != nil { return nil, err }
	pool := make(Variable, length, length)
	for i := int16(0); i < length; i++ {
		left, err := expression(ba, variables)
		if err != nil { return nil, err }
		right, err := expression(ba, variables)
		if err != nil { return nil, err }
		pool[i] = &VariableElement{ name: left.(*Identifier), value: &right }
	}
	return &pool, nil
}
