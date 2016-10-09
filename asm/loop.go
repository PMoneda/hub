package asm

import (
	"github.com/PMoneda/hub/ast"
	"github.com/PMoneda/hub/utils"
)

//LoopCompiler to compile for EXP {...}
type LoopCompiler struct {
}

//Compile FOR statement
func (compiler *LoopCompiler) Compile(offset string, expOffset string, exitOffset string, node *ast.Tree) {

	var cmp ExpCompiler
	Program.Push(expOffset + ":")
	cmp.Compile(node.Value.(utils.Stack))
	Program.Push(JE + " #true :" + offset + " :" + exitOffset)

}
