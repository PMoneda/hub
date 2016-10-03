//Package interpreter execute hub commands
package interpreter

import (
	"errors"
	"fmt"

	"github.com/PMoneda/hub/ast"
	"github.com/PMoneda/hub/lang"
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
		if token == "\n" {
			continue
		}
		interpreter.stmt(&execRoot, token)
		interpreter.root.AppendChild(execRoot)
	}

}

//Print ast
func (interpreter *Interpreter) Print() {
	interpreter.root.Print(func(value interface{}) {
		switch v := value.(type) {
		case lang.Object:
			fmt.Println(v.ToString())
			break
		default:
			fmt.Println(v)
		}

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
func (interpreter *Interpreter) matchKeyword(next string, keyword string) (string, error) {
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

func (interpreter *Interpreter) matchEOL(token string) {
	if token == "\n" {
		return
	}
	panic("End of line expected")
}

//Statement rule for list of commands
func (interpreter *Interpreter) varStmt(root *ast.Tree) {
	var decl ast.DeclVar
	var iden ast.Tree
	var exp ast.Tree
	interpreter.identStmt(&iden)
	next := interpreter.lexer.Next()
	_, err1 := interpreter.matchKeyword(next, "=")
	if err1 != nil {
		panic(err1.Error())
	}
	decl.Op = "="

	interpreter.exprStmt(&exp)
	token := interpreter.lexer.Next()
	interpreter.matchEOL(token)
	root.Value = decl
	root.AppendChild(iden)
	root.AppendChild(exp)
}
func (interpreter *Interpreter) identStmt(parent *ast.Tree) {
	id, err := interpreter.matchIdent()
	if err != nil {
		panic(err.Error())
	}
	var ident = lang.BuildPointer(id)
	parent.Value = ident

}

func (interpreter *Interpreter) exprStmt(root *ast.Tree) {
	interpreter.E(root)
}

//BuildObject based on token
func (interpreter *Interpreter) BuildObject(token string) lang.Object {
	lex := interpreter.lexer
	if lex.IsNumber(token) {
		return lang.BuildNumber(token)
	} else if lex.IsString(token) {
		return lang.BuildString(token)
	}
	return nil
}

//E reflects a expression rule
func (interpreter *Interpreter) E(root *ast.Tree) lang.Object {
	lex := interpreter.lexer
	token := lex.Next()
	if token == "\n" {
		return nil
	} else if lex.IsOperator(token) {
		return lang.BuildOperator(token)
	}
	obj := interpreter.E(root)

	if obj != nil && obj.GetType() == "Operator" {
		root.Value = obj
		var left ast.Tree
		left.Value = interpreter.BuildObject(token)
		root.AppendChild(left)
		var right ast.Tree
		rr := interpreter.E(&right)
		root.AppendChild(right)
		return rr
	} else if obj == nil && !lex.IsOperator(token) {
		t := interpreter.BuildObject(token)
		root.Value = t
	}
	return nil
}
