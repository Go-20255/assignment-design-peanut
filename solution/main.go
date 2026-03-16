package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Number float64
type Bool bool
type Symbol string
type Expr any
type List = []Expr

func parse(chars string) (List, error) {
	idx := 0

	tokens := tokenize(chars)

	exprs := List{}

	for idx < len(tokens) {
		expr, err := parse_tokens(tokens, &idx)

		if err != nil {
			return nil, err
		}

		exprs = append(exprs, expr)
	}

	if idx != len(tokens) {
		return nil, fmt.Errorf("Encountered extra tokens")
	}

	return exprs, nil
}

func parse_tokens(tokens []string, idx *int) (Expr, error) {
	if *idx >= len(tokens) {
		return nil, fmt.Errorf("Encountered EOF early")
	}

	curr := tokens[*idx]
	*idx += 1

	switch curr {
	case "(":
		inner_list := List{}
		for *idx < len(tokens) && tokens[*idx] != ")" {
			parsed_expr, err := parse_tokens(tokens, idx)

			if err != nil {
				return nil, err
			}

			inner_list = append(inner_list, parsed_expr)
		}

		if *idx >= len(tokens) {
			return nil, fmt.Errorf("Encountered unmatched '('")
		}

		*idx += 1
		return inner_list, nil
	case ")":
		return nil, fmt.Errorf("Encountered unexpected ')'")
	default:
		if val, ok := strconv.ParseFloat(curr, 64); ok == nil {
			return Number(val), nil
		} else if curr == "#t" || curr == "#f" {
			return Bool(curr == "#t"), nil
		} else {
			return Symbol(curr), nil
		}
	}
}

func tokenize(chars string) []string {
	paren_expand := strings.Replace(strings.Replace(chars, "(", " ( ", -1), ")", " ) ", -1)
	return strings.Fields(paren_expand)
}

func main() {
	tests := []string{
		// good tests
		"5.4",
		"name",
		"(define name #t)",
		"(a b (c d e))",
		"(begin (define r 10) (* pi (* r r)))",
		"(+ 5 5) (- x -10)",
		// bad tests
		"(()",
		"())",
		"(a",
	}

	for _, test := range tests {
		fmt.Println(parse(test))
	}
}
