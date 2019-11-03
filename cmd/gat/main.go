package main

import (
	"log"
	"os"

	"github.com/donkomura/go-automaton/model"
)

var G *model.Graph = &model.Graph{}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: dfa [binary expression]")
	}
	input := os.Args[1]

	G.InitGraph("sample.dot", "dfa_sample")
	state := G.InitLabel
	var err error
	for _, s := range input {
		state, err = G.Trans(state, string(s))
		if err != nil {
			log.Fatal(err)
		}
	}

	if G.IsFinState(state) {
		log.Print("Accepted.")
	} else {
		log.Print("Invalid")
	}
}
