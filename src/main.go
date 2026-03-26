package main

// TODO: Implement the REPL (Read-Eval-Print-Loop) function
// This function should:
// 1. Print "Beginning REPL. Type \"quit\" to exit."
// 2. Loop forever:
//   - Print "> " as a prompt
//   - Read a line of input
//   - If input is "quit\n", exit the loop
//   - Otherwise, parse the input (using parse())
//   - Evaluate the parsed expression (using eval())
//   - Print the result (if not nil)
//   - Handle parse and semantic errors gracefully
func repl(env Env) {
	// YOUR CODE HERE
}

// TODO: Implement the file() function
// This function should:
// 1. Read the file contents from the given filename
// 2. Tokenize all content at once
// 3. For each non-empty line in the file:
//   - Print "Input: <line>"
//   - Parse one expression starting from the current token index
//   - Evaluate the parsed expression
//   - If result is not nil, print "Output: <result>"
//   - Update the token index to track position
//
// 4. Return any parse or evaluation errors
func file(filename string, env Env) error {
	// YOUR CODE HERE
	return nil
}

// TODO: Implement main()
// This function should:
// 1. Create the initial environment with get_starting_env()
// 2. Check if command-line arguments were provided:
//   - If yes (len(os.Args) > 1): call file() with the filename argument
//   - If no: call repl() to start interactive mode
//
// 3. If file() returns an error, print it and return
func main() {
	// YOUR CODE HERE
}
