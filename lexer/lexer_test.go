package lexer

import "testing"

const (
	tokenize = "./test/tokenize.hub"
)

func TestLexer(t *testing.T) {
	lexer := Lexer{FileName: tokenize}
	for lexer.HasNext() {
		lexer.Next()
	}
}
func TestIsString(t *testing.T) {
	lexer := Lexer{FileName: tokenize}
	isStr := lexer.IsString("\"Hello World\"")
	if !isStr {
		t.Fail()
	}

	isStr1 := lexer.IsString("213")

	if isStr1 {
		t.Fail()
	}
}

func TestIsNumber(t *testing.T) {
	lexer := Lexer{FileName: tokenize}
	isNum := lexer.IsNumber("\"Hello World\"")
	if isNum {
		t.Fail()
	}
	isNum1 := lexer.IsNumber("123")
	if !isNum1 {
		t.Fail()
	}

	isNum2 := lexer.IsNumber("123.123")
	if !isNum2 {
		t.Fail()
	}

	isNum3 := lexer.IsNumber(".123")
	if isNum3 {
		t.Fail()
	}

}
