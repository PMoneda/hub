//Package ast contains the Abstract Syntax Tree Implementation
package ast

//Node is the basic interface for AST
type Node interface {
	TokenList() []string
}

//Expr is a Expresion Node
type Expr interface {
	Node
	exprNode()
}

//Stmt is a Expresion Node
type Stmt interface {
	Node
	stmtNode()
}

//Decl is a Expresion Node
type Decl interface {
	Node
	declNode()
}

////////////////////////////////////////////////////////////////////////////////
//Declare fiedls

////////////////////////////////////////////////////////////////////////////////
