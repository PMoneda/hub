package opcodes

import "fmt"

//Halt represents opcode halt
type Halt struct {
}

//ToString print cpush command
func (opcode Halt) ToString() string {
	return "halt"
}

//Execute execute cpush command
func (opcode Halt) Execute() {
	fmt.Println("Execute Halt")
}
