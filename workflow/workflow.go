// Package workflow controls workflow interpreter
package workflow

import (
	"github.com/PMoneda/hub/asm"
	"github.com/PMoneda/hub/interpreter"
	"github.com/PMoneda/hub/lexer"
)

//Workflow contains information about interpreter
type Workflow struct {
	lex   lexer.Lexer
	inter interpreter.Interpreter
}

// Lex starts to process hub file
func (workflow *Workflow) Lex(fileName string) *Workflow {
	workflow.lex = lexer.Lexer{FileName: fileName}
	workflow.lex.Parse()
	return workflow
}

// BuildAst create ast from file
func (workflow *Workflow) BuildAst() *Workflow {
	workflow.inter = interpreter.Interpreter{}
	workflow.inter.Run(&workflow.lex)
	return workflow
}

// Print print ast tree
func (workflow *Workflow) Print() *Workflow {
	workflow.inter.Print()
	return workflow
}

// Compile ast into hub code
func (workflow *Workflow) Compile() *Workflow {
	asm.Compile(workflow.inter.GetAst())
	return workflow
}
