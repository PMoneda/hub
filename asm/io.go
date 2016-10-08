package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/utils"
)

//PrintCompiler generate print instruction
type PrintCompiler struct {
}

//Compile print statement
func (compiler *PrintCompiler) Compile(expr utils.Stack) {
	var exp ExpCompiler
	exp.Compile(expr)
	Program.Push(PRINT + " r0")
}

//ReadCompiler generate print instruction
type ReadCompiler struct {
}

//Compile print statement
func (compiler *ReadCompiler) Compile(ident lang.Pointer) {
	Program.Push(READ + " " + ident.ToString())
}
