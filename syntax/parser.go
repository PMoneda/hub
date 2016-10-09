//Package syntax assembly the syntax tree
package syntax

import (
	"errors"
	"fmt"

	"github.com/PMoneda/hub/ast"
	"github.com/PMoneda/hub/lang"
	"github.com/PMoneda/hub/lexer"
	"github.com/PMoneda/hub/utils"
)

//Parser instance to run hub code
type Parser struct {
	lexer          *lexer.Lexer
	executionState int
	root           ast.Tree
}

//GetAst built from file
func (parser *Parser) GetAst() *ast.Tree {
	return &parser.root
}

//Run execute a hub script
func (parser *Parser) Run(lexer *lexer.Lexer) {
	parser.lexer = lexer
	parser.root = *new(ast.Tree)
	var begin ast.Begin
	begin.Op = "BEGIN"
	parser.root.Value = begin
	for lexer.HasNext() {
		var execRoot ast.Tree

		token := lexer.Next()
		if lexer.IsBlockDelimiter(token) {
			token = lexer.Next()
		}
		if token == "\\EOF\\" {
			break
		}
		parser.stmt(&execRoot, token)
		parser.root.AppendChild(&execRoot)
	}

}

//Print ast
func (parser *Parser) Print() {
	parser.print(&parser.root)
}

//Print ast
func (parser *Parser) print(root *ast.Tree) {
	root.Print(func(value interface{}) {
		switch v := value.(type) {
		case lang.Object:
			fmt.Println(v.ToString())
			break
		default:
			fmt.Println(v)
		}

	})
}

//exec is the initial point of execution and starts with de first token on hub script
func (parser *Parser) stmt(parent *ast.Tree, token string) error {
	switch token {
	case "var":
		parser.varStmt(parent)
		return nil
	case "print":
		parser.printStmt(parent)
		return nil
	case "read":
		parser.readStmt(parent)
		return nil
	case "if":
		parser.ifStmt(parent)
		return nil
	case "for":
		parser.forStmt(parent)
		return nil
	default:
		if parser.lexer.IsIdent(token) {
			ident := token
			token = parser.lexer.Next()
			if token == "++" || token == "--" {
				var op ast.Tree
				if token == "++" {
					parent.Value = ast.IncStmt{Op: token}
				} else {
					parent.Value = ast.DecStmt{Op: token}
				}
				op.Value = parser.BuildObject(ident)
				token = parser.lexer.Next()
				parser.matchEndOfCommand(token)
				parent.AppendChild(&op)
				return nil
			} else if token == "=" {
				parser.lexer.GiveTokenBack()
				parser.lexer.GiveTokenBack()
				parser.varStmt(parent)
				return nil
			}

		}
		return errors.New("end")

	}
}
func (parser *Parser) printExecInfo() {
	tok := parser.lexer.GetTokenInfo()
	fmt.Printf("error at line %d token %d near %s\n", tok.Line, tok.Index, tok.Value)
}
func (parser *Parser) matchKeyword(next string, keyword string) (string, error) {
	if next != keyword {
		parser.printExecInfo()
		err := errors.New("Expected: " + keyword + " got: " + next)
		//return "", err
		panic(err.Error())
	}
	return next, nil //OK
}

func (parser *Parser) matchIdent() (string, error) {
	next := parser.lexer.Next()
	if !parser.lexer.IsIdent(next) {
		parser.printExecInfo()
		err := errors.New("Expected Identifier  got: " + next)
		return "", err
	}
	return next, nil //OK
}
func (parser *Parser) matchEOF(token string) error {
	if token == "\\EOF\\" {
		tok := parser.lexer.Previous()
		err := errors.New("Unexpected end of file line: " + string(tok.Line))
		return err
	}
	return nil
}

func (parser *Parser) matchEndOfCommand(token string) {
	if token == ";" {
		return
	}
	parser.printExecInfo()
	panic("; expected")
}

func (parser *Parser) matchEOL(token string) {
	if token == "\n" {
		return
	}
	parser.printExecInfo()
	panic("End of line expected")
}

func (parser *Parser) readStmt(root *ast.Tree) {
	var readStmt ast.ReadStmt
	readStmt.Op = "read"
	root.Value = readStmt
	var idenNode ast.Tree
	parser.identStmt(&idenNode)
	next := parser.lexer.Next()
	parser.matchEndOfCommand(next)
	root.AppendChild(&idenNode)
}
func (parser *Parser) stmtBlock(root *ast.Tree, token string) {
	parser.matchKeyword("{", token)
	for token != "}" {
		var left ast.Tree
		token = parser.lexer.Next()

		stmt := parser.stmt(&left, token)
		if stmt == nil {
			root.AppendChild(&left)
		}
	}
	token = parser.lexer.Current()
	parser.matchKeyword("}", token)
}

func (parser *Parser) forStmt(root *ast.Tree) {
	var forStmt ast.ForStmt
	forStmt.Op = "for"

	var blockStmt ast.Block
	blockStmt.Op = "BLOCK_FOR"
	var block ast.Tree

	block.Value = blockStmt

	root.Value = forStmt
	var cond ast.Tree
	parser.exprStmt(&cond)
	token := parser.lexer.Current()
	parser.stmtBlock(&block, token)
	root.AppendChild(&cond)
	root.AppendChild(&block)
}

func (parser *Parser) ifStmt(root *ast.Tree) {
	var ifStmt ast.IfStmt
	ifStmt.Op = "if"
	root.Value = ifStmt
	var cond ast.Tree
	var istrue ast.Tree
	var isElse ast.Tree
	istrue.Value = ast.Block{Op: "IF_BLOCK"}
	isElse.Value = ast.Block{Op: "ELSE_BLOCK"}
	parser.exprStmt(&cond)
	token := parser.lexer.Current()
	parser.stmtBlock(&istrue, token)
	token = parser.lexer.Next()
	if token == "else" {
		token = parser.lexer.Next()
		if token == "if" {
			var right ast.Tree
			parser.ifStmt(&right)
			isElse.AppendChild(&right)
		} else {
			parser.stmtBlock(&isElse, token)
		}
	} else {
		parser.lexer.GiveTokenBack()

	}
	root.AppendChild(&cond)
	root.AppendChild(&istrue)
	if isElse.HasChildren() {
		root.AppendChild(&isElse)
	}

}

func (parser *Parser) printStmt(root *ast.Tree) {
	var printStmt ast.PrintStmt
	printStmt.Op = "print"
	root.Value = printStmt
	var expNode ast.Tree
	parser.exprStmt(&expNode)
	token := parser.lexer.Current()
	parser.matchEndOfCommand(token)
	root.AppendChild(&expNode)
}

//Statement rule for list of commands
func (parser *Parser) varStmt(root *ast.Tree) {
	var decl ast.DeclVar
	var iden ast.Tree
	var exp ast.Tree
	parser.identStmt(&iden)
	next := parser.lexer.Next()
	_, err1 := parser.matchKeyword(next, "=")
	if err1 != nil {
		panic(err1.Error())
	}
	decl.Op = "="
	parser.exprStmt(&exp)
	next = parser.lexer.Current()
	parser.matchEndOfCommand(next)
	root.Value = decl
	root.AppendChild(&iden)
	root.AppendChild(&exp)
}
func (parser *Parser) identStmt(parent *ast.Tree) {
	id, err := parser.matchIdent()
	if err != nil {
		panic(err.Error())
	}
	var ident = lang.BuildPointer(id)
	parent.Value = ident

}

func (parser *Parser) exprStmt(root *ast.Tree) {

	parser.E(root)
}

//BuildObject based on token
func (parser *Parser) BuildObject(token string) lang.Object {
	lex := parser.lexer
	if lex.IsNumber(token) {
		return lang.BuildNumber(token)
	} else if lex.IsString(token) {
		return lang.BuildString(token)
	} else if lex.IsIdent(token) {
		return lang.BuildPointer(token)
	} else if lex.IsOperator(token) {
		return lang.BuildOperator(token)
	} else if lex.IsBoolean(token) {
		return lang.BuildBoolean(token)
	}
	return nil
}

//E reflects a expression rule
func (parser *Parser) E(root *ast.Tree) lang.Object {
	exp := parser.convExpToStack()
	root.Value = exp
	//evaluate the expression
	return nil
}

//Converts a expression to stack of operations
func (parser *Parser) convExpToStack() utils.Stack {
	lex := parser.lexer
	token := lex.Next()
	var terms utils.Stack
	var ops utils.Stack
	for !lex.IsBlockDelimiter(token) && !lex.IsCommandDelimiter(token) && token != "\\EOF\\" {
		if !lex.IsOperator(token) && !lex.IsParenhesis(token) {
			terms.Push(parser.BuildObject(token))
		} else if token == "(" {
			ops.Push(token)
		} else if token == ")" {
			top := ops.Pop()
			for top != "(" {
				terms.Push(top)
				if ops.IsEmpty() {
					parser.printExecInfo()
					panic("Invalid Expression: missing '('")
				} else {
					top = ops.Pop()
				}
			}
		} else {
			currOp := parser.BuildObject(token).(lang.Op)
			if !ops.IsEmpty() {
				switch v := ops.Top().(type) {
				case lang.Object:
					op := v.(lang.Op)
					if op.HighPriority(currOp) {
						terms.Push(ops.Pop())
					}
					break
				}
			}
			ops.Push(currOp)
		}
		token = lex.Next()
	}
	for !ops.IsEmpty() {
		switch v := ops.Pop().(type) {
		case lang.Object:
			terms.Push(v)
			break
		default:
			parser.printExecInfo()
			panic(fmt.Sprintf("Invalid expression  at: %v", v))
		}

	}
	return terms
}
