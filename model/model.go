package model

type Graph struct {
	InitLabel string
	Nodes     []*Node
	FinLabels []string
}

type Node struct {
	From   string
	To     string
	Direct string
}

type DFA interface {
	Add(*Node)
	SetFinLabel(*Node)
	Delete(*Node)
	Trans(*Node, string) *Node
}
