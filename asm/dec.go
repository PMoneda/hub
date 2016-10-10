package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/opcodes"
)

//DecCompiler to compile var Ident--
type DecCompiler struct {
}

//Compile Decrment statement
func (compiler *DecCompiler) Compile(iden lang.Pointer) {
	LoadOp(iden, "r0")
	Program.Push(opcodes.Sub{Op1: iden, Op2: 1, Result: "r0"})
	LoadFromReg("r0", iden)
}
