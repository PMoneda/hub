package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/opcodes"
)

//IncCompiler to compile var Ident++
type IncCompiler struct {
}

//Compile Increment statement
func (compiler *IncCompiler) Compile(iden lang.Pointer) {
	LoadOp(iden, "r0")
	Program.Push(opcodes.Sum{Op1: "1", Op2: "r0", Result: "r0"})
	LoadFromReg("r0", iden)
}
