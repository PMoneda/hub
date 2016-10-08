package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/utils"
)

//ExpCompiler compiles expressions
type ExpCompiler struct {
}

//Compile exp stmt
func (exp *ExpCompiler) Compile(expr utils.Stack) {

	for _, op := range expr {
		switch v := op.(type) {
		case lang.Operator:
			if v.GetSymbol() == "+" {
				Program.Push(POP + " r0")
				Program.Push(POP + " r1")
				Program.Push(SUM + " r0 r1 r0")
				Program.Push(PUSH + " r0")
			}
			if v.GetSymbol() == "*" {
				Program.Push(POP + " r0")
				Program.Push(POP + " r1")
				Program.Push(MULT + " r0 r1 r0")
				Program.Push(PUSH + " r0")
			}
			if v.GetSymbol() == "/" {
				Program.Push(POP + " r0")
				Program.Push(POP + " r1")
				Program.Push(DIV + " r0 r1 r0")
				Program.Push(PUSH + " r0")
			}
			if v.GetSymbol() == "-" {
				Program.Push(POP + " r0")
				Program.Push(POP + " r1")
				Program.Push(SUB + " r0 r1 r0")
				Program.Push(PUSH + " r0")
			}
			break
		case lang.Object:
			PushOp(v)
			break
		}
	}
	Program.Push(POP + " r0")

}
