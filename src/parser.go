package main

import (
	"fmt"
	"strings"
)

// parse() is the top-level parsing function
// It tokenizes the input and calls parse_tokens(), then validates
// that all tokens were consumed
func parse(chars string) (Expr, error) {
	idx := 0

	tokens := tokenize(chars)

	expr, err := parse_tokens(tokens, &idx)

	if err != nil {
		return nil, err
	}

	if idx != len(tokens) {
		return nil, fmt.Errorf("Encountered extra tokens")
	}

	return expr, nil
}

// TODO: Implement parse_tokens()
// This is the core recursive parsing function.
// It takes a slice of tokens and a pointer to the current index.
// It should:
//
//  1. Check if idx is out of bounds -> return error
//  2. Get the current token at tokens[*idx] and increment *idx
//  3. Handle three cases:
//     a) Token is "(" -> parse a list
//     - Create an empty List
//     - Loop while idx < len(tokens) and tokens[*idx] != ")"
//     - Recursively call parse_tokens to parse each element
//     - Append each parsed element to the list
//     - Check if we hit EOF (remaining ")" not found) -> error
//     - Increment idx past the ")"
//     - Return the list
//     b) Token is ")" -> return error (unexpected ")")
//     c) Token is something else -> parse it as a literal
//     - Try parsing as a float using strconv.ParseFloat()
//     - If that works, return Number(val)
//     - Otherwise, check if it's "#t" or "#f"
//     - If so, return Bool(true/false)
//     - Otherwise, return Symbol(token)
func parse_tokens(tokens []string, idx *int) (Expr, error) {
	// YOUR CODE HERE
	return nil, fmt.Errorf("parse_tokens not implemented")
}

// tokenize() converts a string into a slice of tokens
// It should split on parentheses and whitespace
// Example: "(+ 1 2)" -> ["(", "+", "1", "2", ")"]
// This is already implemented for you!
func tokenize(chars string) []string {
	paren_expand := strings.Replace(strings.Replace(chars, "(", " ( ", -1), ")", " ) ", -1)
	return strings.Fields(paren_expand)
}
