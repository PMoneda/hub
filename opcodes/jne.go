package opcodes

import (
	"fmt"
	"strconv"

	"github.com/PMoneda/hub/lang"
)

//Jne represents opcode Jne (Jump Not Equal)
//Jne $a :labelOk :labelNOK Ex: jump to the :label offset if $a != r0
type Jne struct {
	Compare interface{}
	Label   string
}

//ToString print Jne command
func (opcode Jne) ToString() string {
	op := "jne"
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
	return op + " :" + opcode.Label
}

//Execute execute Jne command
func (opcode Jne) Execute() {
	fmt.Println("Execute Jne")
}

//SetOffset execute Je command
func (opcode *Jne) SetOffset(offsets map[string]int) {
	addr1 := offsets[opcode.Label]
	opcode.setOffset(addr1)
}

func (opcode *Jne) setOffset(addr int) {
	opcode.Label = strconv.FormatInt(int64(addr), 10)
}
