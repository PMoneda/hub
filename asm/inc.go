package asm

import "github.com/PMoneda/hub/lang"

//IncCompiler to compile var Ident++
type IncCompiler struct {
}

//Compile Increment statement
func (compiler *IncCompiler) Compile(iden lang.Pointer) {
	Program.Push(MOV + " $" + iden.ToString() + " r0")
	Program.Push(SUM + " #1 r0 r0")
	Program.Push(MOV + " r0 $" + iden.ToString())
}
