package visitors

import (
	"github.com/PMoneda/hub/asm"
	"github.com/PMoneda/hub/ast"
)

//BeginVisitor is the root Visitor in compiler time
type BeginVisitor struct {
}

//Visit visit de first node to create the program
func (visitor *BeginVisitor) Visit(root *ast.Tree) {
	var beginCompiler asm.BeginCompiler
	beginCompiler.Compile(root.Value.(ast.Begin))
	for _, ch := range root.Children {
		var stmtVisitor StmtVisitor
		stmtVisitor.Visit(ch)
	}
	beginCompiler.Halt()
}
