package model

func (g *Graph) CreateInit(node *Node) {
	g.InitNode = node
}

func (g *Graph) Add(node *Node) {
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) CreateFins(node *Node) {
	g.FinNodes = append(g.Nodes, node)
}

func (g *Graph) Delte(node *Node) {
	delete(g.Nodes, find(g.Nodes, node.Current, node.Direct))
}

func (g *Graph) Trans(state Label, token Token) Label {
	return g.Nodes[find(g.Nodes, state, token)].Next
}

func NewNode(current, next, token string) *Node {
	return &Node{
		Current: Label(current),
		Next:    Label(next),
		Direct:  Token(token),
	}
}

func delete(s []*Node, i int) []*Node {
	s = append(s[:i], s[i+1:]...)
	n := make([]*Node, len(s))
	copy(n, s)
	return n
}

func find(s []*Node, c Label, t Token) int {
	for i, node := range s {
		if node.Current == c && node.Direct == t {
			return i
		}
	}
	return 0
}
