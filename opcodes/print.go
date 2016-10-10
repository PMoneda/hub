package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Print is the opcode to print in stdio
//print r0
type Print struct {
	Op interface{}
}

//ToString prints label opcode
func (opcode Print) ToString() string {
	switch v := opcode.Op.(type) {
	case lang.Pointer:
		return "print $" + v.ToString()
	case lang.Number:
		return "print #" + v.ToString()
	case string:
		return "print " + v
	default:
		return "print"
	}
}

//Execute execute label opcode
func (opcode Print) Execute() {
	fmt.Println("Execute Print")
}
