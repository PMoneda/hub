package lexer

import (
	"fmt"
	"testing"
)

const (
	test1 = "./test/test1.hub"
)

func TestLexer(t *testing.T) {
	var lexer Lexer
	lexer.fileName = test1
	lexer.Parse(func(tokens []string) {
		for i := 0; i < len(tokens); i++ {
			fmt.Print(string(tokens[i]))
			fmt.Print("   ")
		}
		fmt.Println()
	})
}

func BenchmarkLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}

}
