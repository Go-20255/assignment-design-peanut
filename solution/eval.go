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
			case "if":
				if len(val) != 4 {
					return nil, fmt.Errorf("Error: 'if' expects 3 arguments (condition, then, else)")
				}
				cond, err := eval(val[1], env)
				if err != nil {
					return nil, err
				}
				condBool, ok := cond.(Bool)
				if !ok {
					return nil, fmt.Errorf("Error: condition in 'if' is not a boolean")
				}
				if condBool {
					return eval(val[2], env)
				} else {
					return eval(val[3], env)
				}
			}
		}

		procedure, err := eval(val[0], env)

		if err != nil {
			return nil, err
		}

		if converted_proc, ok := procedure.(func(List) (any, error)); ok {
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
		}

		if lambda, ok := procedure.(Lambda); ok {
			if len(val[1:]) != len(lambda.params) {
				return nil, fmt.Errorf("Error: lambda expects %d arguments, got %d", len(lambda.params), len(val[1:]))
			}
			callEnv := Env{}
			for k, v := range lambda.env {
				callEnv[k] = v
			}
			for i, param := range lambda.params {
				argVal, err := eval(val[i+1], env)
				if err != nil {
					return nil, err
				}
				callEnv[string(param)] = argVal
			}
			return eval(lambda.body, callEnv)
		}

		return nil, fmt.Errorf("Error: first argument of list is not a procedure")
	default:
		return nil, fmt.Errorf("Error: expr did not match expected cases")
	}
}
