package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/opcodes"
	"github.com/PMoneda/hub/utils"
)

//DeclVarCompiler to compile var Ident = EXP
type DeclVarCompiler struct {
}

//Compile DeclVar statement
func (compiler *DeclVarCompiler) Compile(iden lang.Pointer, exp utils.Stack) {
	Program.Push(opcodes.Load{Op: iden})
	var expc ExpCompiler
	expc.Compile(exp)
	LoadFromReg("r0", iden)

}
