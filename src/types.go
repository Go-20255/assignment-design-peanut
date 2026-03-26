package main

type Number float64
type Bool bool
type Symbol string
type Expr any
type List = []Expr
type Env = map[string]any

// Lambda represents a user-defined procedure
type Lambda struct {
	params []Symbol
	body   Expr
	env    Env
}
