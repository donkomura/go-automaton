package main

import (
	"log"
	"os"

	"github.com/donkomura/go-automaton/model"
)

var G *model.Graph = &model.Graph{}

func Init() {
	G.CreateInit(model.NewNode("", "", ""))
	G.Add(model.NewNode("", "A", "0"))
	G.Add(model.NewNode("", "A", "1"))
	G.Add(model.NewNode("A", "A", "0"))
	G.Add(model.NewNode("A", "B", "1"))
	G.Add(model.NewNode("B", "C", "1"))
	G.Add(model.NewNode("B", "B", "0"))
	G.Add(model.NewNode("C", "C", "0"))
	G.Add(model.NewNode("C", "C", "1"))
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

	if state == model.Label("C") {
		log.Print("Accepted.")
	} else {
		log.Print("Invalid")
	}
}
