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
// This is the core recursive parsing function that converts a token stream into an AST.
// Handle different token types (parentheses, literals, symbols) appropriately.
// See ASSIGNMENT.md section 4.2 and 4.3 for guidance.
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
