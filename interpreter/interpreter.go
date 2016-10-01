//Package interpreter execute hub commands
package interpreter

import (
	"errors"

	"github.com/PMoneda/hub/lexer"
)

//Interpreter instance to run hub code
type Interpreter struct {
	lexer          *lexer.Lexer
	executionState int
}

//Run execute a hub script
func (interpreter *Interpreter) Run(lexer *lexer.Lexer) {
	interpreter.lexer = lexer
	interpreter.stmt()
}

func (interpreter *Interpreter) getNextToken() string {
	if interpreter.lexer.HasNext() {
		return interpreter.lexer.Next()
	}
	return "\\EOF\\"
}

//exec is the initial point of execution and starts with de first token on hub script
func (interpreter *Interpreter) stmt() {
	token := interpreter.getNextToken()
	switch token {
	case "var":
		interpreter.varStmt()
		return
	}
}

func (interpreter *Interpreter) matchKeyword(keyword string) error {
	next := interpreter.getNextToken()
	err := interpreter.matchEOF(next)
	if err == nil {
		if next != keyword {
			err := errors.New("Expected: " + keyword + " got: " + next + " line: " + string(interpreter.lexer.GetCurrentLine()))
			return err
		}
		return nil //OK
	}
	return err
}

func (interpreter *Interpreter) matchIdent() error {
	next := interpreter.getNextToken()
	err := interpreter.matchEOF(next)
	if err == nil {
		if !interpreter.lexer.IsIdent(next) {
			err := errors.New("Expected Identifier  got: " + next + " line: " + string(interpreter.lexer.GetCurrentLine()))
			return err
		}
		return nil //OK
	}
	return err
}
func (interpreter *Interpreter) matchEOF(token string) error {
	if token == "\\EOF\\" {
		err := errors.New("Unexpected end of file line: " + string(interpreter.lexer.GetCurrentLine()))
		return err
	}
	return nil
}

//Statement rule for list of commands
func (interpreter *Interpreter) varStmt() {
	//ID = EXP
	err := interpreter.matchIdent()
	if err != nil {
		panic(err.Error())
	}
	err1 := interpreter.matchKeyword("=")
	if err1 != nil {
		panic(err1.Error())
	}
	err2 := interpreter.expressions()
	if err2 != nil {
		panic(err2.Error())
	}
}

func (interpreter *Interpreter) expressions() error {
	return errors.New("Not Implemented")
}
