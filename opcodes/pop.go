package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Pop represents opcode Pop
//Pop r0 Ex: Pop content to r0
type Pop struct {
	Op interface{}
}

//ToString print Pop command
func (opcode Pop) ToString() string {
	op := "pop "
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

//Execute execute Pop command
func (opcode Pop) Execute() {
	fmt.Println("Execute Pop")
}
