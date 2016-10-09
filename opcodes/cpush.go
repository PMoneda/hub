package opcodes

import "fmt"

//CPush represents opcode cpush
type CPush struct {
}

//ToString print cpush command
func (opcode CPush) ToString() string {
	return "cpush"
}

//Execute execute cpush command
func (opcode CPush) Execute() {
	fmt.Println("Execute CPUSH")
}
