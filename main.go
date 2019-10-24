package main

import (
	"log"
	"os"
)

type Tuple struct {
	Label string
	Bin   string
}

var m map[Tuple]string

func Init() {
	m = make(map[Tuple]string)
	m[Tuple{"", "0"}] = "A"
	m[Tuple{"", "1"}] = "A"
	m[Tuple{"A", "0"}] = "A"
	m[Tuple{"A", "1"}] = "B"
	m[Tuple{"B", "0"}] = "C"
	m[Tuple{"B", "1"}] = "B"
	m[Tuple{"C", "0"}] = "C"
	m[Tuple{"C", "1"}] = "C"
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: dfa [binary expression]")
	}
	input := os.Args[1]

	Init()
	state := ""
	for _, ch := range input {
		state = m[Tuple{state, string(ch)}]
	}

	if state == "C" {
		log.Print("Accepted.")
	} else {
		log.Print("Invalid")
	}
}
