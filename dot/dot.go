package dot

import (
	"fmt"
	"strings"

	"gonum.org/v1/gonum/graph/formats/dot"
	"gonum.org/v1/gonum/graph/formats/dot/ast"
)

type Arrows struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Direct string `json:"direct,omitempty"`
}

func GetGraph(file, name string) (*ast.Graph, error) {
	d, err := dot.ParseFile(file)
	if err != nil {
		return nil, err
	}
	for _, g := range d.Graphs {
		if g.ID == name {
			return g, nil
		}
	}
	return nil, fmt.Errorf("could not find a disignated graph")
}

func GatherArrows(g *ast.Graph) []*Arrows {
	var arrows []*Arrows
	for _, stmt := range g.Stmts {
		if edge, ok := stmt.(*ast.EdgeStmt); ok {
			var directs []string
			if len(edge.Attrs) != 0 {
				directs = strings.Split(edge.Attrs[0].Val, ",")
			}
			if len(directs) == 0 {
				arrows = append(arrows, &Arrows{
					From: edge.From.String(),
					To:   edge.To.Vertex.String(),
				})
			}
			for _, di := range directs {
				di = strings.TrimSpace(strings.Trim(di, "\""))
				arrows = append(arrows, &Arrows{
					From:   edge.From.String(),
					To:     edge.To.Vertex.String(),
					Direct: di,
				})
			}
		}
	}
	return arrows
}

func GetFinLabels(g *ast.Graph) []string {
	var res []string
	for _, stmt := range g.Stmts {
		if nodeStmt, ok := stmt.(*ast.NodeStmt); ok {
			for _, attr := range nodeStmt.Attrs {
				if attr.Val == "doublecircle" {
					res = append(res, nodeStmt.Node.ID)
				}
			}
		}
	}
	return res
}
