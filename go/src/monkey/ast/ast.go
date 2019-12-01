// ast/ast.go

package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

// Statements: do not create values
// e.g. let x = 5;
type Statement interface {
	Node
	statementNode()
}

// Expressions: create values
// e.g. add(5, 10)
type Expression interface {
	Node
	expressionNode()
}

// Program: Root Node of AST
type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // indentifier of binding
	Value Expression  // expression producing a value
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type ReturnStatement struct {
	Token       token.Token // the token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}
