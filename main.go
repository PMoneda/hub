//Package main is the entry point to Hub Interpreter
package main

import (
	"fmt"

	"github.com/PMoneda/hub/interpreter"
	"github.com/PMoneda/hub/lexer"
)

const (
	test1 = "./interpreter/test/test1.hub"
)

func main() {
	fmt.Println("Hello Hub")

	lexer := lexer.Lexer{FileName: test1}
	/*for lexer.HasNext() {
		fmt.Print(lexer.Next() + " ")
	}*/
	var inter interpreter.Interpreter
	inter.Run(&lexer)
	inter.Print()
}
