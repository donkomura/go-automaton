package model

import "github.com/pkg/errors"

func (g *Graph) Add(node *Node) {
	g.Nodes = append(g.Nodes, node)
}

func (g *Graph) SetFinLabel(l string) {
	g.FinLabels = append(g.FinLabels, l)
}

func (g *Graph) Delte(node *Node) error {
	idx, err := find(g.Nodes, node.From, node.Direct)
	if err != nil {
		return err
	}
	delete(g.Nodes, idx)
	return nil
}

func (g *Graph) Trans(s string, t string) (string, error) {
	idx, err := find(g.Nodes, s, t)
	if err != nil {
		return "", err
	}
	return g.Nodes[idx].To, nil
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

func find(s []*Node, c string, t string) (int, error) {
	for i, node := range s {
		if node.From == c && node.Direct == t {
			return i, nil
		}
	}
	return 0, errors.Errorf("not found state direction")
}
