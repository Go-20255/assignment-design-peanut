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

func eval(expr Expr, env Env) (any, error) {
	switch val := expr.(type) {
	case Symbol:
		if sym, ok := env[string(val)]; !ok {
			return nil, fmt.Errorf("Error: symbol was not found in dictionary")
		} else {
			return sym, nil
		}
	case Number, Bool:
		return val, nil
	case List:
		if first, ok := val[0].(Symbol); ok {
			switch first {
			case "define":
				sym, ok := val[1].(Symbol)
				value, err := eval(val[2], env)

				if !ok {
					return nil, fmt.Errorf("Error: left side of define was not a symbol")
				}

				if err != nil {
					return nil, err
				}

				env[string(sym)] = value

				return nil, nil
			}
		}

		procedure, err := eval(val[0], env)

		if err != nil {
			return nil, err
		}

		converted_proc, ok := procedure.(func(List) (any, error))

		if !ok {
			return nil, fmt.Errorf("Error: first argument of list is not a procedure")
		}

		args := List{}
		for _, arg := range val[1:] {
			eval_arg, err := eval(arg, env)

			if err != nil {
				return nil, err
			}

			args = append(args, eval_arg)
		}

		ret, err := converted_proc(args)

		if err != nil {
			return nil, err
		}

		return ret, nil
	default:
		return nil, fmt.Errorf("Error: expr did not match expected cases")
	}
}

func main() {
	tests := []string{
		// good tests
		"5.4",
		"name",
		"(define name #t)",
		"(a b (c d e))",
		"(begin (define r 10) (* pi (* r r)))",
		// bad tests
		"(()",
		"())",
		"(a",
		"(+ 5 5) (- x -10)",
	}

	for _, test := range tests {
		fmt.Println(parse(test))
	}

	env := get_starting_env()

	parsed, _ := parse("(define x 5)")
	res, e := eval(parsed, env)
	fmt.Println(res, e)
	parsed, _ = parse("(+ x 10.8)")
	res, e = eval(parsed, env)
	fmt.Println(res, e)
}
