package main

import (
	"log"
	"os"

	"github.com/donkomura/go-automaton/dot"
	"github.com/donkomura/go-automaton/model"
)

var G *model.Graph = &model.Graph{}

func Init() error {
	graph, err := dot.GetGraph("sample.dot", "dfa_sample")
	if err != nil {
		return err
	}
	for _, l := range dot.GetFinLabels(graph) {
		G.SetFinLabel(l)
	}
	arrows := dot.GatherArrows(graph)
	for _, arrow := range arrows {
		if arrow.Direct == "" {
			G.InitLabel = arrow.To
		}
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
	state := G.InitLabel
	for _, s := range input {
		state = G.Trans(state, string(s))
	}

	if G.IsFinState(state) {
		log.Print("Accepted.")
	} else {
		log.Print("Invalid")
	}
}
