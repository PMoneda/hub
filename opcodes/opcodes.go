package opcodes

//OpCode represents opcode for hub
type OpCode interface {
	ToString() string
	Execute()
}

//FlowControl for the commands that can change execution flow
type FlowControl interface {
	SetOffset(map[string]int)
}
