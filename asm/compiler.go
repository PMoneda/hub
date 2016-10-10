//Package asm is the compiler package for hublang
package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/opcodes"
	"github.com/PMoneda/hub/utils"
)

//Assembler is the assembler process context
type Assembler struct {
	stack             utils.Stack
	offsetMap         map[string]int
	toTranslateOffset []int
}

//Push OpCode to the program
func (asm *Assembler) Push(op interface{}) {
	switch v := op.(type) {
	case opcodes.Label:
		if asm.offsetMap == nil {
			asm.offsetMap = make(map[string]int)
		}
		offset := asm.Len() + 1
		asm.offsetMap[v.Label] = offset
		break
	case opcodes.FlowControl:
		if asm.toTranslateOffset == nil {
			asm.toTranslateOffset = make([]int, 0)
		}
		asm.toTranslateOffset = append(asm.toTranslateOffset, len(asm.stack))
		break
	}
	asm.stack.Push(op)
}

//Len of Program
func (asm *Assembler) Len() int {
	return len(asm.stack)
}

//OffsetMap of labels
func (asm *Assembler) OffsetMap() map[string]int {
	return asm.offsetMap
}

//TranslateOffset of labels
func (asm *Assembler) TranslateOffset() []int {
	return asm.toTranslateOffset
}

//GetStack of Program
func (asm *Assembler) GetStack() *utils.Stack {
	return &asm.stack
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
