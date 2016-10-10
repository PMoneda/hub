package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Read is the opcode to read from stdin
//Read r0
type Read struct {
	Op interface{}
}

//ToString Reads label opcode
func (opcode Read) ToString() string {
	switch v := opcode.Op.(type) {
	case lang.Pointer:
		return "read $" + v.ToString()
	case lang.Number:
		return "read #" + v.ToString()
	case string:
		return "read " + v
	default:
		return "read"
	}
}

//Execute execute label opcode
func (opcode Read) Execute() {
	fmt.Println("Execute Read")
}
