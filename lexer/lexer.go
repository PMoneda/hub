// Package lexer to parse tokens
package lexer

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

//Lexer struct to parse tokens
type Lexer struct {
	file     *os.File
	offset   int
	fileName string
	scanner  *bufio.Scanner
}

//NextLine of code
func (lexer *Lexer) NextLine() string {
	if lexer.file == nil {
		lexer.openFile()
	}

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

//Tokenize line command
func (lexer *Lexer) Tokenize(line string) {

	for i := 0; i < len(line); i++ {
		if isComment(line[i]) {
			return
		} else if line[i] == '"' {
			fmt.Print(string(line[i]))
			i++
			for i < len(line) && line[i] != '"' {
				fmt.Print(string(line[i]))
				i++
			}
			fmt.Print(string(line[i]))
			fmt.Println()
		} else if isIdent(line[i]) {
			for i < len(line) && isIdent(line[i]) {
				fmt.Print(string(line[i]))
				i++
			}
			fmt.Println()
			i--
		} else if isOperator(line[i]) {
			for i < len(line) && isOperator(line[i]) {
				fmt.Print(string(line[i]))
				i++
			}
			fmt.Println()
			i--
		} else if isDelimiter(line[i]) {
			fmt.Println(string(line[i]))
		}

	}
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
