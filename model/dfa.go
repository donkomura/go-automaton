package model

func (g *Graph) Add(node *Node) {
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) SetFinLabel(l string) {
	g.FinLabels = append(g.FinLabels, l)
}

func (g *Graph) Delte(node *Node) {
	delete(g.Nodes, find(g.Nodes, node.From, node.Direct))
}

func (g *Graph) Trans(state string, string string) string {
	return g.Nodes[find(g.Nodes, state, string)].To
}

func NewNode(current, next, direct string) *Node {
	return &Node{
		From:   current,
		To:     next,
		Direct: direct,
	}
}

func (g *Graph) IsFinState(s string) bool {
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

func find(s []*Node, c string, t string) int {
	for i, node := range s {
		if node.From == c && node.Direct == t {
			return i
		}
	}
	return 0
}
