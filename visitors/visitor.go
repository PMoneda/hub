package visitors

import "github.com/PMoneda/hub/ast"

type Visitor interface {
	Visit(*ast.Tree)
}


