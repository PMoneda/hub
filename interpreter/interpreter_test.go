package interpreter

import (
	"testing"

	"github.com/PMoneda/hub/lexer"
)

const (
	test1 = "./test/test1.hub"
)

func TestInterpreter(t *testing.T) {
	lexer := lexer.Lexer{FileName: test1}
	var inter Interpreter
	inter.Run(&lexer)
	inter.Print()

}
