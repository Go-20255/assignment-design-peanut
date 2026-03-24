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
		fmt.Println("DEBUG: Entered file() function")
		os.Stdout.Sync()
	       content, err := os.ReadFile(filename)
	       if err != nil {
		       return fmt.Errorf("Failed to open file: %s", err)
	       }
		fmt.Printf("DEBUG: file content = %q\n", string(content))
		os.Stdout.Sync()

	       tokens := tokenize(string(content))
		fmt.Printf("DEBUG: tokens = %#v\n", tokens)
		os.Stdout.Sync()
	       idx := 0
	       for idx < len(tokens) {
		       parsed, err := parse_tokens(tokens, &idx)
			   fmt.Printf("DEBUG: idx = %d, len(tokens) = %d\n", idx, len(tokens))
			   os.Stdout.Sync()
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
	       return nil
}

func main() {
		f, _ := os.Create("debug_main.txt")
		f.WriteString("DEBUG: main() entered\n")
		f.Close()
		env := get_starting_env()
		fmt.Printf("DEBUG: os.Args = %#v\n", os.Args)
		os.Stdout.Sync()
	       if len(os.Args) > 1 {
		       fmt.Println("DEBUG: Entering file() branch in main")
		       os.Stdout.Sync()
		       filename := os.Args[1]
		       err := file(filename, env)

		       if err != nil {
			       fmt.Printf("Error encountered while running file.\n%s\n", err)
			       return
		       }
	       } else {
		       fmt.Println("DEBUG: Entering repl() branch in main")
		       os.Stdout.Sync()
		       repl(env)
	       }

}
