package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Load is the load variavle opcode
//load $a load new ident on context
type Load struct {
	Op lang.Pointer
}

//ToString prints label opcode
func (opcode Load) ToString() string {
	return "load $" + opcode.Op.ToString()
}

//Execute execute label opcode
func (opcode Load) Execute() {
	fmt.Println("Execute Load")
}
