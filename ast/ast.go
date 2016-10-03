//Package ast contains the Abstract Syntax Tree Implementation
package ast

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

//PrintStmt is a stdio print command
type PrintStmt struct {
	Op string
}

//ReadStmt is a stdin read command
type ReadStmt struct {
	Op string
}
