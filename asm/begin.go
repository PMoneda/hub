package asm

import "github.com/PMoneda/hub/ast"

//BeginCompiler compile entry point
type BeginCompiler struct {
}

//Compile Begin Node
func (begin *BeginCompiler) Compile(ast.Begin) {
	Program.Push("begin:")
	Program.Push("cpush")
}

//Halt program
func (begin *BeginCompiler) Halt() {
	Program.Push("halt")
}
