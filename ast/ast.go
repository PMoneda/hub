//Package ast contains the Abstract Syntax Tree Implementation
package ast

//Node is the basic interface for AST
type Node interface {
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

//PrintStmt is a stdio print command
type PrintStmt struct {
	Op string
}

//ReadStmt is a stdin read command
type ReadStmt struct {
	Op string
}

//IfStmt is if cond Comand
type IfStmt struct {
	Op string
}

//ForStmt is loop  Comand
type ForStmt struct {
	Op string
}

//Begin is entry point node
type Begin struct {
	Op string
}

//Block is a block of code
type Block struct {
	Op string
}
