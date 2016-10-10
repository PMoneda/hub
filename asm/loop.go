package asm

import (
	"github.com/PMoneda/hub/ast"
	"github.com/PMoneda/hub/opcodes"
	"github.com/PMoneda/hub/utils"
)

//LoopCompiler to compile for EXP {...}
type LoopCompiler struct {
}

//Compile FOR statement
func (compiler *LoopCompiler) Compile(offset string, expOffset string, exitOffset string, node *ast.Tree) {
	var cmp ExpCompiler
	Program.Push(opcodes.Label{Label: expOffset})
	cmp.Compile(node.Value.(utils.Stack))
	Program.Push(opcodes.Je{Compare: true, LabelOk: offset, LabelNOk: exitOffset})

}
