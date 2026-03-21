package main

import "fmt"

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
