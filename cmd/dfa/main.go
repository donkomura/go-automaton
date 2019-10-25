package main

import (
	"log"
	"os"

	"github.com/donkomura/go-automaton/dot"
	"github.com/donkomura/go-automaton/model"
)

var G *model.Graph = &model.Graph{}

func Init() error {
	G.CreateInit(model.NewNode("", "", ""))
	graph, err := dot.NewGraph("sample.dot", "dfa_sample")
	if err != nil {
		return err
	}
	arrows := dot.GatherArrows(graph)
	for _, arrow := range arrows {
		G.Add(model.NewNode(arrow.From, arrow.To, arrow.Direct))
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: dfa [binary expression]")
	}
	input := os.Args[1]

	Init()
	state := G.InitNode.Current
	for _, s := range input {
		state = G.Trans(state, model.Token(s))
	}

	if state == model.Label("A") {
		log.Print("Accepted.")
	} else {
		log.Print("Invalid")
	}
}