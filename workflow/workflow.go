// Package workflow controls workflow interpreter
package workflow

import (
	"fmt"

	"github.com/PMoneda/hub/asm"
	"github.com/PMoneda/hub/interpreter"
	"github.com/PMoneda/hub/lexer"
	"github.com/PMoneda/hub/visitors"
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
	var begin visitors.BeginVisitor
	begin.Visit(workflow.inter.GetAst())
	for i, cmd := range asm.Program {
		fmt.Print(fmt.Sprintf("0x%08d", i))
		fmt.Print("   ")
		fmt.Print(cmd)
		fmt.Println()
	}
	return workflow
}
