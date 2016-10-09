package opcodes

import "fmt"

//Label is the offset opcode
type Label struct {
	Label string
}

//ToString prints label opcode
func (label Label) ToString() string {
	return label.Label + ":"
}

//Execute execute label opcode
func (label Label) Execute() {
	fmt.Println("Execute Label")
}
