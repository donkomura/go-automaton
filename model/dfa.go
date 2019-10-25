package model

func (g *Graph) Add(node *Node) {
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) CreateFins(l Label) {
	g.FinLabels = append(g.FinLabels, l)
}

func (g *Graph) Delte(node *Node) {
	delete(g.Nodes, find(g.Nodes, node.From, node.Direct))
}

func (g *Graph) Trans(state Label, token Token) Label {
	return g.Nodes[find(g.Nodes, state, token)].To
}

func NewNode(current, next, token string) *Node {
	return &Node{
		From:   Label(current),
		To:     Label(next),
		Direct: Token(token),
	}
}

func (g *Graph) IsFinState(s Label) bool {
	for _, fin := range g.FinLabels {
		if fin == s {
			return true
		}
	}
	return false
}

func delete(s []*Node, i int) []*Node {
	s = append(s[:i], s[i+1:]...)
	n := make([]*Node, len(s))
	copy(n, s)
	return n
}

func find(s []*Node, c Label, t Token) int {
	for i, node := range s {
		if node.From == c && node.Direct == t {
			return i
		}
	}
	return 0
}
