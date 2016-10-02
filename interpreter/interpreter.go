//Package interpreter execute hub commands
package interpreter

import (
	"errors"
	"fmt"

	"github.com/PMoneda/hub/ast"
	"github.com/PMoneda/hub/lexer"
)

//Interpreter instance to run hub code
type Interpreter struct {
	lexer          *lexer.Lexer
	executionState int
	root           ast.Tree
}

//Run execute a hub script
func (interpreter *Interpreter) Run(lexer *lexer.Lexer) {
	interpreter.lexer = lexer
	interpreter.root = *new(ast.Tree)
	interpreter.root.Value = "BEGIN"
	for lexer.HasNext() {
		var execRoot ast.Tree
		token := lexer.Next()
		if token == "\\EOF\\" {
			break
		}
		interpreter.stmt(&execRoot, token)
		interpreter.root.AppendChild(execRoot)
	}

}

//Print ast
func (interpreter *Interpreter) Print() {
	interpreter.root.DeepWalk(func(value interface{}) {
		fmt.Println(value)
	})
}

//exec is the initial point of execution and starts with de first token on hub script
func (interpreter *Interpreter) stmt(parent *ast.Tree, token string) {
	switch token {
	case "var":
		interpreter.varStmt(parent)
		return
	}
}
func (interpreter *Interpreter) printExecInfo() {
	fmt.Print("line: ")
	fmt.Print(interpreter.lexer.GetCurrentLine())
	fmt.Print(" token: ")
	fmt.Print(interpreter.lexer.GetCurrentToken())
	fmt.Println()
	fmt.Println("Line")
	fmt.Println(interpreter.lexer.GetLine())
}
func (interpreter *Interpreter) matchKeyword(keyword string) (string, error) {
	next := interpreter.lexer.Next()
	if next != keyword {
		interpreter.printExecInfo()
		err := errors.New("Expected: " + keyword + " got: " + next)
		return "", err
	}
	return next, nil //OK
}

func (interpreter *Interpreter) matchIdent() (string, error) {
	next := interpreter.lexer.Next()
	if !interpreter.lexer.IsIdent(next) {
		interpreter.printExecInfo()
		err := errors.New("Expected Identifier  got: " + next)
		return "", err
	}
	return next, nil //OK
}
func (interpreter *Interpreter) matchEOF(token string) error {
	if token == "\\EOF\\" {
		err := errors.New("Unexpected end of file line: " + string(interpreter.lexer.GetCurrentLine()))
		return err
	}
	return nil
}

//Statement rule for list of commands
func (interpreter *Interpreter) varStmt(root *ast.Tree) {
	var decl ast.DeclVar
	var iden ast.Tree
	var exp ast.Tree
	interpreter.identStmt(&iden)

	_, err1 := interpreter.matchKeyword("=")
	if err1 != nil {
		panic(err1.Error())
	}
	decl.Op = "="

	interpreter.exprStmt(&exp)

	root.Value = decl
	root.AppendChild(iden)
	root.AppendChild(exp)
}
func (interpreter *Interpreter) identStmt(parent *ast.Tree) {
	id, err := interpreter.matchIdent()
	if err != nil {
		panic(err.Error())
	}
	var ident = ast.Ident{Name: id}
	parent.Value = ident

}

func (interpreter *Interpreter) exprStmt(root *ast.Tree) {
	token := interpreter.lexer.Next()
	root.Value = token

}
