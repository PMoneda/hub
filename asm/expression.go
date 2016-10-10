package asm

import (
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/opcodes"
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
				Program.Push(opcodes.Pop{Op: "r0"})
				Program.Push(opcodes.Inv{Op1: "r0", Result: "r0"})
				Program.Push(opcodes.Push{Op: "r0"})
				break
			}
			Program.Push(opcodes.Pop{Op: "r0"})
			Program.Push(opcodes.Pop{Op: "r1"})
			if v.GetSymbol() == "+" {
				Program.Push(opcodes.Sum{Op1: "r0", Op2: "r1", Result: "r0"})
			} else if v.GetSymbol() == "*" {
				Program.Push(opcodes.Mult{Op1: "r0", Op2: "r1", Result: "r0"})
			} else if v.GetSymbol() == "/" {
				Program.Push(opcodes.Div{Op1: "r0", Op2: "r1", Result: "r0"})
			} else if v.GetSymbol() == "**" {
				Program.Push(opcodes.Pow{Op1: "r0", Op2: "r1", Result: "r0"})
			} else if v.GetSymbol() == "-" {
				Program.Push(opcodes.Sub{Op1: "r0", Op2: "r1", Result: "r0"})
			} else if v.GetSymbol() == ">" {
				Program.Push(opcodes.Gt{Op1: "r1", Op2: "r0", Result: "r0"})
			} else if v.GetSymbol() == ">=" {
				Program.Push(opcodes.Gte{Op1: "r1", Op2: "r0", Result: "r0"})

			} else if v.GetSymbol() == "<" {
				Program.Push(opcodes.Lt{Op1: "r1", Op2: "r0", Result: "r0"})
			} else if v.GetSymbol() == "<=" {
				Program.Push(opcodes.Lte{Op1: "r1", Op2: "r0", Result: "r0"})
			} else if v.GetSymbol() == "==" {
				Program.Push(opcodes.Eq{Op1: "r1", Op2: "r0", Result: "r0"})
			} else if v.GetSymbol() == "!=" {
				Program.Push(opcodes.Diff{Op1: "r1", Op2: "r0", Result: "r0"})
			}
			Program.Push(opcodes.Push{Op: "r0"})
			break
		case lang.Object:
			PushOp(v)
			break
		}
	}
	Program.Push(opcodes.Pop{Op: "r0"})

}
