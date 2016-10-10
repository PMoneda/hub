package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Je represents opcode Je (Jump Equal)
//je $a :labelOk :labelNOK Ex: jump to the :label offset if $a == r0
type Je struct {
	Compare  interface{}
	LabelOk  string
	LabelNOk string
}

//ToString print Je command
func (opcode Je) ToString() string {
	op := "je"
	switch v := opcode.Compare.(type) {
	case lang.Pointer:
		op += " $" + v.ToString()
		break
	case lang.Object:
		op += " #" + v.ToString()
		break
	default:
		op += fmt.Sprintf(" #%v", v)
		break
	}
	return op + " :" + opcode.LabelOk + " :" + opcode.LabelNOk
}

//Execute execute Je command
func (opcode Je) Execute() {
	fmt.Println("Execute Je")
}
