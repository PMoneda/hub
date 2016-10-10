package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Inv represents opcode Inv
//Inv #true r0 Ex: execute r0 = !true
type Inv struct {
	Op1    interface{}
	Result interface{}
}

//ToString print Inv command
func (opcode Inv) ToString() string {
	op := "Inv"
	switch v := opcode.Op1.(type) {
	case lang.Pointer:
		op += " $" + v.ToString()
		break
	case lang.Object:
		op += " #" + v.ToString()
		break
	default:
		op += fmt.Sprintf(" %v", v)
		break
	}

	switch v := opcode.Result.(type) {
	case lang.Pointer:
		op += " $" + v.ToString()
		break
	case lang.Object:
		op += " #" + v.ToString()
		break
	case string:
		op += fmt.Sprintf(" %v", v)
		break
	default:
		op += fmt.Sprintf(" #%v", v)
		break
	}

	return op
}

//Execute execute Inv command
func (opcode Inv) Execute() {
	fmt.Println("Execute Inv")
}
