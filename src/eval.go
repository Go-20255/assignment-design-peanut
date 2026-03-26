package main

import "fmt"

// TODO: Implement eval()
// This is the core evaluation function that processes parsed expressions.
// It takes an expression and an environment (map of variable bindings) and returns:
// - (any, error): the result of evaluation, or an error if something goes wrong
//
// Implementation strategy:
// 1. Use type assertion: switch val := expr.(type)
// 2. Handle 4 cases:
//
// CASE 1: Symbol
//   - Look up the symbol in the environment (env[string(val)])
//   - If not found, return error "symbol was not found in dictionary"
//   - If found, return the value
//
// CASE 2: Number, Bool
//   - These are self-evaluating
//   - Simply return them as-is
//
// CASE 3: List (the complex case)
//
//   - First, check if the first element is a special form (define, if, lambda)
//     To do this: assert val[0] is a Symbol and check its name
//
//   - SPECIAL FORM: "define"
//     Syntax: (define symbol value)
//
//   - Ensure val[1] is a Symbol (the variable name)
//
//   - Evaluate val[2] to get the value
//
//   - Store in environment: env[string(symbol)] = value
//
//   - Return nil, nil (define returns nil)
//
//   - SPECIAL FORM: "if"
//     Syntax: (if condition then-expr else-expr)
//
//   - Ensure exactly 4 elements (if + 3 args)
//
//   - Evaluate the condition (val[1])
//
//   - Assert it's a Bool
//
//   - If true, evaluate and return val[2]
//
//   - If false, evaluate and return val[3]
//
//   - SPECIAL FORM: "lambda"
//     Syntax: (lambda (param1 param2 ...) body-expr)
//
//   - Ensure exactly 3 elements
//
//   - Assert val[1] is a List (the parameter list)
//
//   - Convert each parameter to a Symbol
//
//   - **IMPORTANT**: Capture the current environment in capturedEnv
//
//   - Return a Lambda struct with params, body, and capturedEnv
//
//   - REGULAR FUNCTION CALL (not a special form):
//     Syntax: (func arg1 arg2 ...)
//
//   - Evaluate val[0] to get the procedure
//
//   - Evaluate all arguments (val[1:])
//
//   - If procedure is a built-in func(List) (any, error):
//
//   - Call it with the evaluated arguments
//
//   - If procedure is a Lambda:
//
//   - Check argument count matches parameter count
//
//   - Create a new Env with the Lambda's captured environment
//
//   - Bind parameters to arguments in the new Env
//
//   - Evaluate the Lambda's body in the new Env
//
//   - Otherwise, return error "first argument of list is not a procedure"
//
// CASE 4: Default
//   - Return error "expr did not match expected cases"
func eval(expr Expr, env Env) (any, error) {
	// YOUR CODE HERE
	return nil, fmt.Errorf("eval not implemented")
}
