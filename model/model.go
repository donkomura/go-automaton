package model

type Label string
type Token string

type Graph struct {
	InitNode *Node
	Nodes    []*Node
	FinNodes []*Node
}

type Node struct {
	Current Label
	Next    Label
	Direct  Token
}

type DFA interface {
	CreateInit(*Node)
	Add(*Node)
	CreateFins(*Node)
	Delete(*Node)
	Trans(*Node, Token) *Node
}
