package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Push represents opcode Push
//push r0 Ex: push content of r0 on stack
type Push struct {
	Op interface{}
}

//ToString print Push command
func (opcode Push) ToString() string {
	op := "push "
	switch v := opcode.Op.(type) {
	case lang.Pointer:
		op += "$" + v.ToString()
		break
	case lang.Object:
		op += "#" + v.ToString()
		break
	default:
		op += fmt.Sprintf("%v", v)
		break
	}
	return op
}

//Execute execute Push command
func (opcode Push) Execute() {
	fmt.Println("Execute Push")
}
