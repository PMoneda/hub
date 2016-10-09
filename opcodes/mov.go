package opcodes

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
)

//Mov represents opcode mov
//mov op1 op2 Ex: Mov content from op1 to op2
type Mov struct {
	From interface{}
	To   interface{}
}

//ToString print mov command
func (opcode Mov) ToString() string {
	op := "mov "
	switch v := opcode.From.(type) {
	case lang.Pointer:
		op += "$" + v.ToString()
		break
	case lang.Number:
		op += "#" + v.ToString()
		break
	case string:
		op += v
		break
	}
	op += " "
	switch v := opcode.To.(type) {
	case lang.Pointer:
		op += "$" + v.ToString()
		break
	case lang.Number:
		op += "#" + v.ToString()
		break
	case string:
		op += v
		break
	}
	return op
}

//Execute execute mov command
func (opcode Mov) Execute() {
	fmt.Println("Execute MOV")
}
