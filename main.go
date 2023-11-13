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

	parsed := parseProgram(content)
	engine := machine.NewEngine(
		parsed,
		100000000,
	)
	if err := engine.Eval(); err != nil {
		log.Fatal(err)
	}
}

func parseProgram(input []byte) []byte {
	output := []byte{}
	for _, v := range input {
		if v == '+' ||
			v == '-' ||
			v == '>' ||
			v == '<' ||
			v == '[' ||
			v == ']' ||
			v == '.' ||
			v == ',' {
			output = append(output, v)
		}
	}
	return output
}
