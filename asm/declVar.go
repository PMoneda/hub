package asm

import (
	"fmt"

	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/utils"
)

//DeclVarCompiler to compile var Ident = EXP
type DeclVarCompiler struct {
}

//Compile DeclVar statement
func (compiler *DeclVarCompiler) Compile(iden lang.Pointer, exp utils.Stack) {
	Program.Push(fmt.Sprintf("%s $%s", LOAD, iden.ToString()))
	var expc ExpCompiler
	expc.Compile(exp)
	Program.Push(fmt.Sprintf("%s r0 $%s", MOV, iden.ToString()))
}
