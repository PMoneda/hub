//Package interpreter execute hub commands
package interpreter

import "github.com/PMoneda/hub/lexer"

//Interpreter instance to run hub code
type Interpreter struct {
	lexer *lexer.Lexer
}

//Program is the initial point of execution and starts with de first token on hub script
func (interpreter *Interpreter) Program(token string) {

}

//Statement rule for list of commands
func (interpreter *Interpreter) Statement(token string) {
	switch token {
	case "var":
		//Execute var statement
		return
	}
}

func isVar(token string) bool {
	return token == "var"
}
func isIF(token string) bool {
	return token == "if"
}
func isElif(token string) bool {
	return token == "elif"
}
func isElse(token string) bool {
	return token == "else"
}
func isReturn(token string) bool {
	return token == "return"
}
func isDataSource(token string) bool {
	return token == "datasource"
}
func isGet(token string) bool {
	return token == "get"
}
func isPost(token string) bool {
	return token == "post"
}
func isImport(token string) bool {
	return token == "import"
}
func isAnd(token string) bool {
	return token == "and"
}
func isOr(token string) bool {
	return token == "or"
}
func isNot(token string) bool {
	return token == "not"
}
func isListen(token string) bool {
	return token == "listen"
}
func isDefun(token string) bool {
	return token == "defun"
}
