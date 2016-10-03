// Package lexer to parse tokens
package lexer

import (
	"bufio"
	"bytes"
	"os"
	"unicode"
	"unicode/utf8"
)

//Lexer struct to parse tokens
type Lexer struct {
	file         *os.File
	offset       int
	FileName     string
	scanner      *bufio.Scanner
	currentLine  int
	tokenBuffer  []string
	currentToken int
	line         string
}

//GetCurrentLine gets current lexer linenumber
func (lexer *Lexer) GetCurrentLine() int {
	return lexer.currentLine
}

//GetLine get current line
func (lexer *Lexer) GetLine() string {
	return lexer.line
}

//GetCurrentToken returns index of current token
func (lexer *Lexer) GetCurrentToken() int {
	return lexer.currentToken
}

//NextLine of code
func (lexer *Lexer) NextLine() string {
	text := lexer.scanner.Text()
	lexer.currentLine++
	lexer.line = text
	return text
}

//Next return the next token
func (lexer *Lexer) Next() string {
	if len(lexer.tokenBuffer) > lexer.currentToken {
		token := lexer.tokenBuffer[lexer.currentToken]
		lexer.currentToken++
		return token
	} else if lexer.HasNext() {
		line := lexer.NextLine()
		tokens := lexer.Tokenize(line)
		for len(tokens) == 0 && lexer.HasNext() {
			line = lexer.NextLine()
			tokens = lexer.Tokenize(line)
		}
		if !lexer.HasNext() {
			return "\\EOF\\"
		}
		lexer.currentToken = 1
		lexer.tokenBuffer = lexer.Tokenize(line)

		return lexer.tokenBuffer[0]
	}

	return "\\EOF\\"
}

func (lexer *Lexer) openFile() {
	f, err := os.Open(lexer.FileName)
	if err != nil {
		panic(err.Error())
	}
	lexer.file = f
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	lexer.scanner = scanner
}

//HasNext line to read
func (lexer *Lexer) HasNext() bool {
	if lexer.scanner == nil {
		lexer.openFile()
		lexer.currentToken = 0
		line := lexer.NextLine()
		lexer.tokenBuffer = lexer.Tokenize(line)
		return true
	}
	if lexer.currentToken < len(lexer.tokenBuffer) {
		return true
	}
	hasNext := lexer.scanner.Scan()
	if hasNext {
		lexer.currentToken = 0
		line := lexer.NextLine()
		lexer.tokenBuffer = lexer.Tokenize(line)
	}
	return hasNext
}

//Tokenize line command
func (lexer *Lexer) Tokenize(line string) []string {
	var tokens []string
	var buffer bytes.Buffer
	if len(line) > 0 && line[0] == '#' {
		return tokens
	}
	for i := 0; i < len(line); i++ {
		if isComment(line[i]) {
			if len(tokens) > 0 {
				tokens = append(tokens, "\n")
			}
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
	tokens = append(tokens, "\n")
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
	case "**":
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
	return c == '.' || 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' || c == '_' || c >= utf8.RuneSelf
}

//IsIdent returns if token is a identificator
func (lexer *Lexer) IsIdent(token string) bool {
	isDel := lexer.IsDelimiter(token)
	isOp := lexer.IsOperator(token)
	isKeyword := lexer.IsKeyword(token)
	isStr := lexer.IsString(token)
	isNum := lexer.IsNumber(token)
	isBool := lexer.IsBoolean(token)
	return !isDel && !isOp && !isKeyword && !isStr && !isNum && !isBool
}

//IsString returns if token is a string
func (lexer *Lexer) IsString(token string) bool {
	return byte(token[0]) == '"' && byte(token[len(token)-1]) == '"'
}

//IsBoolean returns if token is a boolean
func (lexer *Lexer) IsBoolean(token string) bool {
	return token == "true" || token == "false"
}

//IsNumber returns if token is a string
func (lexer *Lexer) IsNumber(token string) bool {
	r, _ := utf8.DecodeRuneInString(token)
	return unicode.IsNumber(r)
}

//IsKeyword returns true if token is reserved word
func (lexer *Lexer) IsKeyword(token string) bool {
	switch token {
	case "var":
		return true
	case "const":
		return true
	case "if":
		return true
	case "for":
		return true
	case "elif":
		return true
	case "else":
		return true
	case "and":
		return true
	case "or":
		return true
	case "not":
		return true
	case "mod":
		return true
	case "defun":
		return true
	case "listen":
		return true
	case "get":
		return true
	case "post":
		return true

	}
	return false
}

//IsAssingOp return true if operator is =
func (lexer *Lexer) IsAssingOp(token string) bool {
	return token == "="
}
