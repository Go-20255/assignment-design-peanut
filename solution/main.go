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
			fmt.Printf("Parse error: failed to parse input. %s\n", err)
			continue
		}

		evaled, err := eval(parsed, env)

		if err != nil {
			fmt.Printf("Semantic error: failed to evaluate input. %s\n", err)
			continue
		}

		fmt.Println(evaled)
	}
}

func main() {
	repl()
}
