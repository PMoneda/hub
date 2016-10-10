package opcodes

import "fmt"

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
