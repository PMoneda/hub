package lexer

import (
	"fmt"
	"testing"
)

const (
	tokenize = "./test/tokenize.hub"
)

func TestLexer(t *testing.T) {
	lexer := Lexer{fileName: tokenize}
	for lexer.HasNext() {
		fmt.Print(lexer.Next() + " ")
	}
}

func BenchmarkLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}

}
