package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Sub represents opcode Sub
//Sub r0 r1 r0 Ex: execute r0 = r0 - r1
type Sub struct {
	Op1    interface{}
	Op2    interface{}
	Result interface{}
}

//ToString print Sub command
func (opcode Sub) ToString() string {
	op := "sub"
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

//Execute execute Sub command
func (opcode Sub) Execute() {
	fmt.Println("Execute Sub")
}
