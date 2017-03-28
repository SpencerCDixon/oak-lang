package ast

import (
	"github.com/spencercdixon/monkey/token"
)

// Every node in the AST must satisfy this interface
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// This is the ROOT node of every AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

//--------------
// Let Statement
//--------------
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // name of our variable
	Value Expression  // expression that the variable gets assigned to
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

//--------------
// Identifier
//--------------
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

//--------------
// Return Statements
//--------------
type ReturnStatement struct {
	Token token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
