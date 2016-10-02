//Package ast contains the Abstract Syntax Tree Implementation
package ast

import "fmt"

//Node is the basic interface for AST
type Node interface {
	Print()
}

//Expr is a Expresion Node
type Expr struct {
}

//Ident is a Identification
type Ident struct {
	Name string
}

//RootNode data strucutre for declarations
type RootNode struct {
	Program string
}

//DeclVar data strucutre for declarations
type DeclVar struct {
	Op string
}

//Print prints AST node
func (dec DeclVar) Print() {
	fmt.Println(dec.Op)
}

//Print prints AST node
func (ident Ident) Print() {
	fmt.Println("Name: " + ident.Name)
}

//Print prints AST node
func (dec Expr) Print() {
	fmt.Println("Ola Mundo")
}

//Print root
func (root RootNode) Print() {
	fmt.Println("Start")
}
