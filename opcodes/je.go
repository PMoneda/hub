package opcodes

import (
	"fmt"
	"strconv"

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

//SetOffset execute Je command
func (opcode *Je) SetOffset(offsets map[string]int) {
	addr1 := offsets[opcode.LabelOk]
	addr2 := offsets[opcode.LabelNOk]
	opcode.setOffset(addr1, addr2)
}

func (opcode *Je) setOffset(addr1 int, addr2 int) {
	opcode.LabelOk = strconv.FormatInt(int64(addr1), 10)
	opcode.LabelNOk = strconv.FormatInt(int64(addr2), 10)
}
