package syntax

import (
	"testing"

	"github.com/PMoneda/hub/lexer"
)

const (
	test1 = "./test/test1.hub"
)

func TestParser(t *testing.T) {
	lexer := lexer.Lexer{FileName: test1}
	var inter Parser
	inter.Run(&lexer)
	inter.Print()

}
