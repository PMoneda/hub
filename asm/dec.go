package asm

import "github.com/PMoneda/hub/lang"

//DecCompiler to compile var Ident--
type DecCompiler struct {
}

//Compile Decrment statement
func (compiler *DecCompiler) Compile(iden lang.Pointer) {
	Program.Push(MOV + " $" + iden.ToString() + " r0")
	Program.Push(SUB + " #1 r0 r0")
	Program.Push(MOV + " r0 $" + iden.ToString())
}
