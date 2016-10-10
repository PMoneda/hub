//Package main is the entry point to Hub parser
package main

import (
	"fmt"

	"github.com/PMoneda/hub/workflow"
)

const (
	test1 = "./syntax/test/test1.hub"
)

func main() {
	fmt.Println("Hello Hub")
	var flow workflow.Workflow
	//flow.Lex(test1).BuildAst().Print().Compile().PrintAsm()
	flow.Lex(test1).BuildAst().Compile().PrintAsm()
}
