package dot

import (
	"fmt"
	"log"
	"strings"

	"gonum.org/v1/gonum/graph/formats/dot"
	"gonum.org/v1/gonum/graph/formats/dot/ast"
)

type Arrows struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Direct string `json:"direct,omitempty"`
}

func NewGraph(file, name string) (*ast.Graph, error) {
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
			for _, di := range directs {
				arrows = append(arrows, &Arrows{
					From:   edge.From.String(),
					To:     edge.To.Vertex.String(),
					Direct: di,
				})
				log.Printf("From: %v, To: %v, Attr: %v\n", edge.From.String(), edge.To.Vertex.String(), di)
			}
		}
	}
	return arrows
}
