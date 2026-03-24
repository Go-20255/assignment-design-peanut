package main

import (
	"fmt"
	"reflect"
)

func add(args List) (any, error) {
	total := 0.0

	for _, expr := range args {
		val, ok := expr.(Number)

		if !ok {
			return 0, fmt.Errorf("Error: attemtped to add non-number")
		}

		total += float64(val)
	}

	return Number(total), nil
}

func mult(args List) (any, error) {
	total := 1.0

	for _, expr := range args {
		val, ok := expr.(Number)

		if !ok {
			return 0, fmt.Errorf("Error: attemtped to multiply non-number")
		}

		total *= float64(val)
	}

	return Number(total), nil
}

func sub(args List) (any, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("Error: no arguments in subtraction")
	}

	val, ok := args[0].(Number)
	if !ok {
		return 0, fmt.Errorf("Error: attemtped to subtract non-number")
	}

	if len(args) == 1 {
		return Number(-val), nil
	}

	total := val

	for _, expr := range args[1:] {
		val, ok := expr.(Number)

		if !ok {
			return 0, fmt.Errorf("Error: attemtped to subtract non-number")
		}

		total -= val
	}

	return total, nil
}

func div(args List) (any, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("Error: no arguments in division")
	}

	val, ok := args[0].(Number)
	if !ok {
		return 0, fmt.Errorf("Error: attemtped to divide non-number")
	}

	if len(args) == 1 {
		return 1 / val, nil
	}

	total := val

	for _, expr := range args[1:] {
		val, ok := expr.(Number)

		if !ok {
			return 0, fmt.Errorf("Error: attemtped to divide non-number")
		}

		total /= val
	}

	return total, nil
}

func and(args List) (any, error) {
	for _, expr := range args {
		val, ok := expr.(Bool)

		if !ok {
			continue
		}

		if val == false {
			return Bool(false), nil
		}
	}

	return Bool(true), nil
}

func or(args List) (any, error) {
	for _, expr := range args {
		val, ok := expr.(Bool)

		if !ok {
			continue
		}

		if val == true {
			return Bool(true), nil
		}
	}

	return Bool(false), nil
}

func not(args List) (any, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("Error: no arguments in not")
	}

	if len(args) != 1 {
		return 0, fmt.Errorf("Error: too many arguments in not")
	}

	val, ok := args[0].(Bool)

	if !ok {
		return Bool(false), nil
	}

	return Bool(val == false), nil
}

func equal(args List) (any, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("Error: no arguments in equal")
	}

	if len(args) == 1 {
		return Bool(true), nil
	}

	prev := args[0]

	for _, expr := range args[1:] {
		if !reflect.DeepEqual(prev, expr) {
			return Bool(false), nil
		}
	}

	return Bool(true), nil
}

func get_starting_env() Env {
	env := Env{}
	env["+"] = add
	env["*"] = mult
	env["-"] = sub
	env["/"] = div
	env["and"] = and
	env["or"] = or
	env["not"] = not
	env["eq?"] = equal

	return env
}
