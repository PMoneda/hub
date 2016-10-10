package visitors

import (
	"fmt"
	"strconv"

	"github.com/PMoneda/hub/asm"
	"github.com/PMoneda/hub/ast"
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/opcodes"
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
		size := asm.Program.Len()
		offset := "if_" + strconv.FormatInt(int64(size), 10)
		elseOffset := "else_" + strconv.FormatInt(int64(size), 10)
		//Compile if condition
		ifStmt.Compile(offset, elseOffset, node)
		ifBlock := node.Children[1]
		var stmtVisitor StmtVisitor
		asm.Program.Push(opcodes.CPush{})
		stmtVisitor.Visit(ifBlock.Children[0])
		asm.Program.Push(opcodes.CPop{})
		asm.Program.Push(opcodes.Jmp{Label: offset})

		if len(node.Children) == 3 {
			//has else block
			elseBlock := node.Children[2]
			asm.Program.Push(elseOffset + ":")
			stmtVisitor.Visit(elseBlock)
			asm.Program.Push(opcodes.Jmp{Label: offset})
		}
		asm.Program.Push(offset + ":")
		break
	case ast.Block:
		var stmtVisitor StmtVisitor
		stmtVisitor.Visit(node.Children[0])
		break
	case ast.ForStmt:
		var forStmt asm.LoopCompiler
		size := asm.Program.Len()
		offset := "for_block_" + strconv.FormatInt(int64(size), 10)
		expOffset := "for_" + strconv.FormatInt(int64(size), 10)
		exitOffset := "exit_for_" + strconv.FormatInt(int64(size), 10)
		exp := node.Children[0]
		block := node.Children[1]
		asm.Program.Push(opcodes.CPush{})
		forStmt.Compile(offset, expOffset, exitOffset, exp)
		var stmtVisitor StmtVisitor
		asm.Program.Push(offset + ":")
		stmtVisitor.Visit(block.Children[0])
		asm.Program.Push(opcodes.Jmp{Label: offset})
		asm.Program.Push(exitOffset + ":")
		asm.Program.Push(opcodes.CPop{})

		break
	case ast.IncStmt:
		var incStmt asm.IncCompiler
		incStmt.Compile(node.Children[0].Value.(lang.Pointer))
		break
	case ast.DecStmt:
		var decStmt asm.DecCompiler
		decStmt.Compile(node.Children[0].Value.(lang.Pointer))
		break
	default:
		fmt.Println(v)

	}
}
