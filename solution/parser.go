package main

import (
	"fmt"
	"strconv"
	"strings"
)

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

func parse_tokens(tokens []string, idx *int) (Expr, error) {
	if *idx >= len(tokens) {
		return nil, fmt.Errorf("Encountered EOF early")
	}

	fmt.Printf("DEBUG: parse_tokens idx=%d token=%#v\n", *idx, tokens[*idx])
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
