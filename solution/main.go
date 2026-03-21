package main

import (
	"fmt"
)

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
