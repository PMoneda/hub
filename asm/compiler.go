//Package asm is the compiler package for hublang
package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/opcodes"
	"github.com/PMoneda/hub/utils"
)

//Assembler is the assembler process context
type Assembler struct {
	stack utils.Stack
}

//Push OpCode to the program
func (asm *Assembler) Push(op interface{}) {
	asm.stack.Push(op)
}

//Len of Program
func (asm *Assembler) Len() int {
	return len(asm.stack)
}

//GetStack of Program
func (asm *Assembler) GetStack() utils.Stack {
	return asm.stack
}

//Program is a Global Assembler object
var Program Assembler

//LoadOp on register
func LoadOp(op1 lang.Object, register string) {
	op := opcodes.Mov{}
	op.From = op1
	op.To = register
	Program.Push(op)
}

//LoadFromReg ident from reg
func LoadFromReg(register string, op1 lang.Object) {
	op := opcodes.Mov{}
	op.From = register
	op.To = op1
	Program.Push(op)
}

//PushOp push operand on stack
func PushOp(op1 lang.Object) {
	op := opcodes.Push{Op: op1}
	Program.Push(op)
}

//PopOp pop operand from stack
func PopOp(op1 lang.Object) {
	op := opcodes.Pop{Op: op1}
	Program.Push(op)
}
