//Package asm is the compiler package for hublang
package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/utils"
)

//OPCODES for hub-lang
const (
	CPUSH = "cpush"
	CPOP  = "cpop"
	MOV   = "mov"
	LOAD  = "load"
	PRINT = "print"
	SUM   = "sum"
	MULT  = "mult"
	DIV   = "div"
	POW   = "pow"
	JMP   = "jmp"
	SUB   = "sub"
	PUSH  = "push"
	POP   = "pop"
)

//OpCode represents opcode for hub
type OpCode struct {
	OpCode string
}

//Program is a list of commands
var Program utils.Stack

//LoadOp on register
func LoadOp(op1 lang.Object, register string) {
	op := MOV
	if op1.GetType() == "Pointer" {
		op += " $" + op1.ToString()
	} else if op1.GetType() == "Number" {
		op += " #" + op1.ToString()
	} else if op1.GetType() == "String" {
		op += " " + op1.ToString()
	}
	op += " " + register
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
