// Package lexer to parse tokens
package lexer

import (
	"bufio"
	"bytes"
	"os"
	"unicode/utf8"
)

//Lexer struct to parse tokens
type Lexer struct {
	file         *os.File
	offset       int
	fileName     string
	scanner      *bufio.Scanner
	currentLine  int
	tokenBuffer  []string
	currentToken int
}

//GetCurrentLine gets current lexer linenumber
func (lexer *Lexer) GetCurrentLine() int {
	return lexer.currentLine
}

//NextLine of code
func (lexer *Lexer) NextLine() string {
	if lexer.file == nil {
		lexer.openFile()
	}
	lexer.currentLine++
	return lexer.scanner.Text()
}

//Next return the next token
func (lexer *Lexer) Next() string {
	if lexer.currentToken >= len(lexer.tokenBuffer) {
		if lexer.HasNext() {
			line := lexer.NextLine()
			lexer.currentToken = 0
			lexer.tokenBuffer = lexer.Tokenize(line)
		}
	}
	token := lexer.tokenBuffer[lexer.currentToken]
	lexer.currentToken++
	return token
}

func (lexer *Lexer) openFile() {
	f, _ := os.Open(lexer.fileName)
	lexer.file = f
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lexer.scanner = scanner
}

//HasNext line to read
func (lexer *Lexer) HasNext() bool {
	if lexer.scanner == nil {
		lexer.openFile()
	}
	if lexer.currentToken < len(lexer.tokenBuffer) {
		return true
	}
	return lexer.scanner.Scan()
}

//Tokenize line command
func (lexer *Lexer) Tokenize(line string) []string {
	tokens := make([]string, 1)
	var buffer bytes.Buffer
	for i := 0; i < len(line); i++ {
		if isComment(line[i]) {
			return tokens
		} else if line[i] == '"' {
			buffer.WriteString(string(line[i]))
			i++
			for i < len(line) && line[i] != '"' {
				buffer.WriteString(string(line[i]))
				i++
			}
			buffer.WriteString(string(line[i]))
			tokens = append(tokens, buffer.String())
		} else if isIdent(line[i]) {
			for i < len(line) && isIdent(line[i]) {
				buffer.WriteString(string(line[i]))
				i++
			}
			i--
			tokens = append(tokens, buffer.String())
		} else if lexer.IsOperator(string(line[i])) {
			for i < len(line) && lexer.IsOperator(string(line[i])) {
				buffer.WriteString(string(line[i]))
				i++
			}
			i--
			tokens = append(tokens, buffer.String())
		} else if lexer.IsDelimiter(string(line[i])) {
			buffer.WriteString(string(line[i]))
			tokens = append(tokens, buffer.String())
		}
		buffer.Reset()
	}
	return tokens
}
func isComment(c byte) bool {
	return c == '#'
}

//IsOperator returns true if token is a operator
func (lexer *Lexer) IsOperator(c string) bool {
	switch c {
	case "<":
		return true
	case ">":
		return true
	case "=":
		return true
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	case ";":
		return true
	case "!":
		return true
	case "?":
		return true

	}

	return false
}

//IsDelimiter returns true with token is a block delimiters
func (lexer *Lexer) IsDelimiter(c string) bool {
	switch c {
	case "(":
		return true
	case ")":
		return true
	case "{":
		return true
	case "}":
		return true
	}

	return false
}

func isIdent(c byte) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || c == '_' || c >= utf8.RuneSelf
}

func (lexer *Lexer) IsIdent(token string) bool {
	return !lexer.IsDelimiter(token) && !lexer.IsOperator(token) && !lexer.IsKeyword(token)
}

//IsKeyword returns true if token is reserved word
func (lexer *Lexer) IsKeyword(token string) bool {
	switch token {
	case "var":
		return false
	case "const":
		return false
	case "if":
		return false
	case "for":
		return false
	case "elif":
		return false
	case "else":
		return false
	case "and":
		return false
	case "or":
		return false
	case "not":
		return false
	case "mod":
		return false
	case "defun":
		return false
	case "listen":
		return false
	case "get":
		return false
	case "post":
		return false

	}
	return true
}

//IsAssingOp return true if operator is =
func (lexer *Lexer) IsAssingOp(token string) bool {
	return token == "="
}
