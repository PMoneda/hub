package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Div represents opcode Div
//Div r0 r1 r0 Ex: execute r0 = r0 / r1
type Div struct {
	Op1    interface{}
	Op2    interface{}
	Result interface{}
}

//ToString print Div command
func (opcode Div) ToString() string {
	op := "div"
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
	switch v := opcode.Op2.(type) {
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

//Execute execute Div command
func (opcode Div) Execute() {
	fmt.Println("Execute Div")
}
