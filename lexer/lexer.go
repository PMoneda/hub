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
	file      *os.File
	offset    int
	fileName  string
	scanner   *bufio.Scanner
	totalLine int
}

//NextLine of code
func (lexer *Lexer) NextLine() string {
	if lexer.file == nil {
		lexer.openFile()
	}
	lexer.totalLine++
	return lexer.scanner.Text()
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
	return lexer.scanner.Scan()
}

//Parse parses line of command
func (lexer *Lexer) Parse(callback func([]string)) {
	for lexer.HasNext() {
		tokens := lexer.Tokenize(lexer.NextLine())
		callback(tokens)
	}

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
		} else if isOperator(line[i]) {
			for i < len(line) && isOperator(line[i]) {
				buffer.WriteString(string(line[i]))
				i++
			}
			i--
			tokens = append(tokens, buffer.String())
		} else if isDelimiter(line[i]) {
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
func isOperator(c byte) bool {
	switch c {
	case '<':
		return true
	case '>':
		return true
	case '=':
		return true
	case '+':
		return true
	case '-':
		return true
	case '*':
		return true
	case '/':
		return true
	case ';':
		return true
	case '!':
		return true
	case '?':
		return true

	}

	return false
}

func isDelimiter(c byte) bool {
	switch c {
	case '(':
		return true
	case ')':
		return true
	case '{':
		return true
	case '}':
		return true
	}

	return false
}

func isIdent(c byte) bool {
	return 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || c == '_' || c >= utf8.RuneSelf
}
