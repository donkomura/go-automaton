package model

type Label string
type Token string

type Graph struct {
	InitLabel Label
	Nodes     []*Node
	FinLabels []Label
}

type Node struct {
	From   Label
	To     Label
	Direct Token
}

type DFA interface {
	CreateInit(*Node)
	Add(*Node)
	CreateFins(*Node)
	Delete(*Node)
	Trans(*Node, Token) *Node
}
