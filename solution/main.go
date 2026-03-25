package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Failed to open file: %s", err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	allTokens := tokenize(string(content))
	tokenIdx := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Parse one complete expression from the line
		expr, err := parse_tokens(allTokens, &tokenIdx)
		if err != nil {
			return fmt.Errorf("Parse error: failed to parse input.\n%s\n", err)
		}

		fmt.Printf("Input: %s\n", line)

		evaled, err := eval(expr, env)
		if err != nil {
			return fmt.Errorf("Semantic error: failed to evaluate input.\n%s\n", err)
		}
		if evaled != nil {
			fmt.Printf("Output: %v\n", evaled)
		}
	}
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
