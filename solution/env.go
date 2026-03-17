package main

import "fmt"

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

func get_starting_env() Env {
	env := Env{}
	env["+"] = add

	return env
}
