//Package asm is the compiler package for hublang
package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/opcodes"
	"github.com/PMoneda/hub/utils"
)

//OPCODES for hub-lang
const (
	PRINT              = "print" //print $a Ex: print content of $a on stdio
	SUM                = "sum"   //sum r0 r1 r0 Ex: execute r0 = r0 + r1
	MULT               = "mult"  //mult r0 r1 r0 Ex: execute r0 = r0 * r1
	DIV                = "div"   //div r0 r1 r0 Ex: execute r0 = r0 / r1
	POW                = "pow"   //pow r0 r1 r0 Ex: execute r0 = r0 ** r1
	JMP                = "jmp"   //jmp :label Ex: jump to the :label offset
	JE                 = "je"    //je $a :label Ex: jump to the :label offset if $a == r0
	JNE                = "jne"   //jne $a :label Ex: jump to the :label offset if $a != r0
	SUB                = "sub"   //sub r0 r1 r0 Ex: execute r0 = r0 - r1
	PUSH               = "push"  //push r0 Ex: push content of r0 on stack
	POP                = "pop"   //pop r0 Ex: pop content from stack and move to r0
	READ               = "read"  //read $a Ex: read stdin and put result on $a
	LessThan           = "lt"    //lt $a #2 Ex: execute r0 = $a < 2 (true or false)
	GreaterThan        = "gt"    //gt $a #2 Ex: execute r0 = $a > 2 (true or false)
	LessThanOrEqual    = "lte"   //lte $a #2 Ex: execute r0 = $a <= 2 (true or false)
	GreaterThanOrEqual = "gte"   //gte $a #2 Ex: execute r0 = $a >= 2 (true or false)
	EQ                 = "eq"    //eq $a #2 Ex: execute r0 = $a == 2 (true or false)
	DIFF               = "diff"  //diff $a #2 Ex: execute r0 = $a != 2 (true or false)
	INV                = "inv"   //inv #true r0 Ex: execute r0 = not true
)

//Program is a list of commands
var Program utils.Stack

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
	op := PUSH
	if op1.GetType() == "Pointer" {
		op += " $" + op1.ToString()
	} else if op1.GetType() == "Number" {
		op += " #" + op1.ToString()
	} else if op1.GetType() == "String" {
		op += " " + op1.ToString()
	}
	Program.Push(op)
}

//PopOp pop operand from stack
func PopOp(op1 lang.Object) {
	op := POP
	if op1.GetType() == "Pointer" {
		op += " $" + op1.ToString()
	} else if op1.GetType() == "Number" {
		op += " #" + op1.ToString()
	} else if op1.GetType() == "String" {
		op += " " + op1.ToString()
	}
	Program.Push(op)
}
