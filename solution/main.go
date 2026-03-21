package main

import (
	"bufio"
	"fmt"
	"os"
)

func repl(env Env) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Beginning REPL. Type \"quit\" to exit.")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')

		if input == "quit\n" {
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

func file(filename string, env Env) error {
	program, err := os.Open(filename)

	if err != nil {
		return fmt.Errorf("Failed to open file: %s", err)
	}

	scanner := bufio.NewScanner(program)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		parsed, err := parse(line)

		if err != nil {
			return fmt.Errorf("Parse error: failed to parse input.\n%s\n", err)
		}

		evaled, err := eval(parsed, env)

		if err != nil {
			return fmt.Errorf("Semantic error: failed to evaluate input.\n%s\n", err)
		}

		if evaled != nil {
			fmt.Println(evaled)
		}
	}

	program.Close()

	return nil
}

func main() {
	env := get_starting_env()

	if len(os.Args) > 1 {
		filename := os.Args[1]
		err := file(filename, env)

		if err != nil {
			fmt.Printf("Error encountered while running file.\n%s\n", err)
			return
		}
	} else {
		repl(env)
	}

}
