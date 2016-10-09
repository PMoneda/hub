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
			if v.GetSymbol() == "not" {
				Program.Push(POP + " r0")
				Program.Push(INV + " r0 r0")
				Program.Push(PUSH + " r0")
				break
			}
			Program.Push(POP + " r0")
			Program.Push(POP + " r1")
			assign := " r0 r1 r0"
			if v.GetSymbol() == "+" {
				Program.Push(SUM + assign)
			} else if v.GetSymbol() == "*" {
				Program.Push(MULT + assign)
			} else if v.GetSymbol() == "/" {
				Program.Push(DIV + assign)
			} else if v.GetSymbol() == "-" {
				Program.Push(SUB + assign)
			} else if v.GetSymbol() == ">" {
				Program.Push(GreaterThan + " r1 r0 r0")
			} else if v.GetSymbol() == ">=" {
				Program.Push(GreaterThanOrEqual + " r1 r0 r0")
			} else if v.GetSymbol() == "<" {
				Program.Push(LessThan + " r1 r0 r0")
			} else if v.GetSymbol() == "<=" {
				Program.Push(LessThanOrEqual + " r1 r0 r0")
			} else if v.GetSymbol() == "==" {
				Program.Push(EQ + assign)
			} else if v.GetSymbol() == "!=" {
				Program.Push(GreaterThan + assign)
			}
			Program.Push(PUSH + " r0")
			break
		case lang.Object:
			PushOp(v)
			break
		}
	}
	Program.Push(POP + " r0")

}
