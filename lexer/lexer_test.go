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
	for lexer.HasNext() {
		lexer.Tokenize(lexer.NextLine())
		fmt.Println()
	}

}

func BenchmarkLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}

}
