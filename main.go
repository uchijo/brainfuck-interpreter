package main

import (
	"flag"
	"log"

	"github.com/uchijo/brainfuck-interpreter/machine"
)

type cmdArgs struct {
	input string
}

func main() {
	flag.Parse()
	rawArgs := flag.Args()
	args := cmdArgs{
		input: rawArgs[0],
	}

	engine := machine.NewEngine(
		args.input,
		1000,
	)
	if err := engine.Eval(); err != nil {
		log.Fatal(err)
	}
}
