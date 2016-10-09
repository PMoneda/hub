package opcodes

import "fmt"

//CPop represents opcode cpop
type CPop struct {
}

//ToString print cpop command
func (opcode CPop) ToString() string {
	return "cpop"
}

//Execute execute cpop command
func (opcode CPop) Execute() {
	fmt.Println("Execute CPOP")
}
