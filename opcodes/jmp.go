package opcodes

import (
	"fmt"
	"strconv"
)

//Jmp represents opcode Jmp
//Jmp r0 r1 r0 Ex: execute r0 = r0 + r1
type Jmp struct {
	Label string
}

//ToString print Jmp command
func (opcode Jmp) ToString() string {
	return "jmp :" + opcode.Label
}

//Execute execute Jmp command
func (opcode Jmp) Execute() {
	fmt.Println("Execute Jmp")
}

//SetOffset execute Je command
func (opcode *Jmp) SetOffset(offsets map[string]int) {
	addr1 := offsets[opcode.Label]
	opcode.setOffset(addr1)

}

func (opcode *Jmp) setOffset(addr int) {
	opcode.Label = strconv.FormatInt(int64(addr), 10)
}
