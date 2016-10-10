package asm

import (
	"github.com/PMoneda/hub/ast"
	"github.com/PMoneda/hub/opcodes"
	"github.com/PMoneda/hub/utils"
)

//IfCompiler to compile if EXP {...}
type IfCompiler struct {
}

//Compile IF statement
func (compiler *IfCompiler) Compile(offset string, elseOffset string, node *ast.Tree) {
	exp := node.Children[0]
	var cmp ExpCompiler
	cmp.Compile(exp.Value.(utils.Stack))
	if len(node.Children) == 3 {
		Program.Push(opcodes.Jne{Compare: true, Label: elseOffset})
	} else {
		Program.Push(opcodes.Jne{Compare: true, Label: offset})
	}

}
