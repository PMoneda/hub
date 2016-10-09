package visitors

import (
	"fmt"
	"strconv"

	"github.com/PMoneda/hub/asm"
	"github.com/PMoneda/hub/ast"
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/utils"
)

//StmtVisitor to statements nodes
type StmtVisitor struct {
}

//Visit stmt node
func (visitor *StmtVisitor) Visit(node *ast.Tree) {
	//Visit statements
	switch v := node.Value.(type) {
	case ast.DeclVar:
		var declStmt asm.DeclVarCompiler
		iden := node.Children[0]
		exp := node.Children[1]
		declStmt.Compile(iden.Value.(lang.Pointer), exp.Value.(utils.Stack))
		break
	case ast.PrintStmt:
		var printStmt asm.PrintCompiler
		exp := node.Children[0]
		printStmt.Compile(exp.Value.(utils.Stack))
		break
	case ast.ReadStmt:
		var readStmt asm.ReadCompiler
		exp := node.Children[0]
		readStmt.Compile(exp.Value.(lang.Pointer))
		break
	case ast.IfStmt:

		var ifStmt asm.IfCompiler
		size := len(asm.Program)
		offset := "if_" + strconv.FormatInt(int64(size), 10)
		ifStmt.Compile(offset, node)
		exp := node.Children[0]
		for _, n := range exp.Children {
			var stmtVisitor StmtVisitor
			stmtVisitor.Visit(n)
		}
		asm.Program.Push(offset + ":")

		break
	default:
		fmt.Println(v)
	}
}
