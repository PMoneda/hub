// Package workflow controls workflow interpreter
package workflow

import "../lexer"

//Workflow contains information about interpreter
type Workflow struct {
}

// Lex starts to process hub file
func (workflow *Workflow) Lex(fileName string) {
	lexer.Parser(fileName)
}
