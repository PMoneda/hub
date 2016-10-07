//Package asm is the compiler package for hublang
package asm

import (
	"fmt"

	"github.com/PMoneda/hub/ast"
)

//OpCode represents opcode for hub
type OpCode struct {
	OpCode string
}

//Compile ast into hub asm
func Compile(ast *ast.Tree) {
	ast.Walk(func(node interface{}) {
		fmt.Print(node)
	})
}
