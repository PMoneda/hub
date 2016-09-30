package lexer

import (
	"fmt"
	"testing"
)

const (
	tokenize = "./test/tokenize.hub"
)

func TestLexer(t *testing.T) {
	var lexer Lexer
	lexer.fileName = tokenize
	lexer.Parse(func(tokens []string) {
		for i := 0; i < len(tokens); i++ {
			fmt.Print(string(tokens[i]))
		}
		fmt.Println()
	})
}

func BenchmarkLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}

}
