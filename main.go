package main

import (
	"log"

	"github.com/uchijo/brainfuck-interpreter/machine"
)

func main() {
	engine := machine.NewEngine(
		"+++++++++[>++++++++>+++++++++++>+++>+<<<<-]>.>++.+++++++..+++.>+++++.<<+++++++++++++++.>.+++.------.--------.>+.>+.",
		1000,
	)
	if err := engine.Eval(); err != nil {
		log.Fatal(err)
	}
}
