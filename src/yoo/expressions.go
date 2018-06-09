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

type Function struct {
	callee Any
	args []Any
}
func NewFunction(ba *ByteArray) (*Function, error) {
	ret, err := ba.ReadInt64()
	if err != nil { return nil, err }
	bytes, err := ba.ReadBytesA(int(ret))
	if err != nil { return nil, err }
	callee, err := expr(bytes)
	if err != nil { return nil, err }
	length, err := bytes.ReadInt16()
	if err != nil { return nil, err }

	pool := make([]Any, length, length)
	for i := int16(0); i < length; i++ {
		arg, err := expr(bytes)
		if err != nil { return nil, err }
		pool[i] = arg
	}
	return &Function{ callee: callee, args: pool }, nil
}
