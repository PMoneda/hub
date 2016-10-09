package opcodes

//OpCode represents opcode for hub
type OpCode interface {
	ToString() string
	Execute()
}