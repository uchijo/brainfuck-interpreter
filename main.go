package main

import (
	"fmt"
	"strings"

	"github.com/uchijo/brainfuck-interpreter/machine"
)

func main() {
	input := "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++."
	env := machine.NewEnv()
	fmt.Println(env)
	eval(input)
}

func eval(input string) {
	env := machine.NewEnv()
	for _, char := range strings.Split(input, "") {
		switch char {
		case "+":
			env.Incr()
		case ".":
			fmt.Printf("%v", env.CurrentString())
		}
	}
}

