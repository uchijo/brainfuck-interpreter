package main

import (
	"flag"
	"log"
	"os"

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

	content, err := os.ReadFile(args.input)
	if err != nil {
		panic(err)
	}

	engine := machine.NewEngine(
		content,
		100000000,
	)
	if err := engine.Eval(); err != nil {
		log.Fatal(err)
	}
}
