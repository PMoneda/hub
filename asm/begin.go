package asm

import (
	"github.com/PMoneda/hub/ast"
	"github.com/PMoneda/hub/opcodes"
)

//BeginCompiler compile entry point
type BeginCompiler struct {
}

//Compile Begin Node
func (begin *BeginCompiler) Compile(ast.Begin) {
	Program.Push(opcodes.Label{Label: "begin"})
	Program.Push(opcodes.CPush{})
}

//Halt program
func (begin *BeginCompiler) Halt() {
	Program.Push(opcodes.Halt{})
}
