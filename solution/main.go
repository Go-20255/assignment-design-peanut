package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Beginning REPL. Type \"quit\" to exit.")

	env := get_starting_env()

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')

		if input == "quit" {
			break
		}

		if err != nil {
			fmt.Println("Error: failed to capture input")
			continue
		}

		parsed, err := parse(input)

		if err != nil {
			fmt.Printf("Parse error: failed to parse input.\n%s\n", err)
			continue
		}

		evaled, err := eval(parsed, env)

		if err != nil {
			fmt.Printf("Semantic error: failed to evaluate input.\n%s\n", err)
			continue
		}

		if evaled != nil {
			fmt.Println(evaled)
		}
	}
}

func main() {
	repl()
}
