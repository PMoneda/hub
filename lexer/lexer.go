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
	file        *os.File
	offset      int
	FileName    string
	nextToken   int
	tokens      []Token
	totalTokens int
}

//Token describe a basic struct for tokens
type Token struct {
	Value string
	Line  int
	Index int
}

//GetTokenInfo return info about current token
func (lexer *Lexer) GetTokenInfo() Token {
	return lexer.tokens[lexer.nextToken-1]
}

//Current token
func (lexer *Lexer) Current() string {
	return lexer.tokens[lexer.nextToken-1].Value
}

//Previous token
func (lexer *Lexer) Previous() Token {
	if lexer.nextToken == 0 {
		return lexer.tokens[0]
	}
	return lexer.tokens[lexer.nextToken-2]
}

//GiveTokenBack to the unread state
func (lexer *Lexer) GiveTokenBack() {
	lexer.nextToken--
	if lexer.nextToken < 0 {
		lexer.nextToken = 0
	}
}

//Next return the next token
func (lexer *Lexer) Next() string {
	if lexer.nextToken >= len(lexer.tokens) {
		return "\\EOF\\"
	}
	curr := lexer.tokens[lexer.nextToken].Value
	lexer.nextToken++
	return curr
}

//Parse read the code file and execute  lexical parsing
func (lexer *Lexer) Parse() {
	f, err := os.Open(lexer.FileName)
	if err != nil {
		panic(err.Error())
	}
	lexer.file = f
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	totalLine := 1
	for scanner.Scan() {
		line := scanner.Text()
		tokenBuffer := lexer.Tokenize(line)
		for i := 0; i < len(tokenBuffer); i++ {
			var tok Token
			tok.Value = tokenBuffer[i]
			tok.Line = totalLine
			tok.Index = i + 1
			lexer.tokens = append(lexer.tokens, tok)
		}
		totalLine++
	}
	lexer.nextToken = 0
	lexer.totalTokens = len(lexer.tokens)
}

//HasNext line to read
func (lexer *Lexer) HasNext() bool {
	return lexer.nextToken < lexer.totalTokens
}

//Tokenize line command
func (lexer *Lexer) Tokenize(line string) []string {
	var tokens []string
	var buffer bytes.Buffer
	for i := 0; i < len(line); i++ {
		if line[i] == '\n' {
			continue
		}
		if isComment(line[i]) {
			return tokens
		} else if line[i] == ';' {
			buffer.WriteString(string(line[i]))
			tokens = append(tokens, buffer.String())
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
	case "and":
		return true
	case "or":
		return true
	case "not":
		return true
	case "mod":
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

//IsParenhesis returns true with token is a block parenthesis
func (lexer *Lexer) IsParenhesis(c string) bool {
	switch c {
	case "(":
		return true
	case ")":
		return true
	}

	return false
}

//IsBlockDelimiter returns true with token is a block delimiter
func (lexer *Lexer) IsBlockDelimiter(c string) bool {
	switch c {
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

//IsCommandDelimiter returns true if token is ;
func (lexer *Lexer) IsCommandDelimiter(token string) bool {
	return token == ";"
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
