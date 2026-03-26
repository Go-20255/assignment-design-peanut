# Lisp Interpreter Assignment - CSCI 541/641

**Build Your Own Lisp Interpreter in Go!**

## Quick Start

```bash
# Build the interpreter
make build

# Run all tests
make test

# Or use the test runner script
./run.sh
```

## What is This?

This is a **one-week programming assignment** for a college Go class where you'll implement a functional Lisp interpreter from scratch. You'll learn about parsing, type systems, environments, and recursive evaluation.

Visit [ASSIGNMENT.md](ASSIGNMENT.md) for complete details, including:
- Learning objectives and requirements
- Implementation guidance for each component
- 10 test cases covering basic through advanced features
- Full rubric (100 points)
- Common mistakes and debugging tips

## Project Structure

```
├── ASSIGNMENT.md          # Full assignment details & rubric
├── Makefile               # Build and test automation
├── run.sh                 # Test runner script
├── src/                   # Student starter code (YOU FILL THIS IN)
│   ├── main.go           # [TODO] Entry point
│   ├── parser.go         # [TODO] Tokenization & parsing
│   ├── eval.go           # [TODO] Expression evaluation
│   ├── types.go          # (Provided)
│   ├── env.go            # (Provided)
│   └── go.mod
├── solution/             # Reference solution (don't peek!)
├── input/                # Test input files (10 tests)
└── output/               # Expected output files
```

## What You Need to Implement

Only **three files** have TODO sections:
- **`src/main.go`**: REPL and file input handling
- **`src/parser.go`**: Tokenization and parsing
- **`src/eval.go`**: Expression evaluation with special forms

## Supported Language Features

Your interpreter will support:
- **Arithmetic**: `+`, `-`, `*`, `/`
- **Comparison**: `eq?` (equality)
- **Variables**: `(define x 7)` then `x`
- **Conditionals**: `(if #t expr1 expr2)`
- **Functions**: `(lambda (x) (+ x 1))`
- **REPL**: Interactive prompt
- **File mode**: Execute code from a file

## Example Usage

```lisp
> (+ 1 2)
3
> (define inc (lambda (x) (+ x 1)))
> (inc 5)
6
> (if #t "yes" "no")
yes
> quit
```

## Learning Outcomes

By the end of this assignment, you'll understand:
- ✓ Recursive descent parsing
- ✓ Abstract syntax tree (AST) construction
- ✓ Type systems and type assertions in Go
- ✓ Environment management and variable scoping
- ✓ Functional programming and closures
- ✓ Error reporting and debugging

## Testing

All 10 test cases provided:
1. Arithmetic operations
2. Variable binding with `define`
3. Conditional expressions with `if`
4. Equality comparison
5. Lambda functions with captures
6. And more!

Run `make test` or `./run.sh` to validate your implementation against expected outputs.

## Rubric Summary

- **Code Quality (60 pts)**: Parser, Evaluator, Special Forms, Advanced features
- **Functionality (25 pts)**: All 10 test cases must pass
- **Style & Documentation (10 pts)**: Readable code and proper error messages
- **Bonus (5 pts)**: REPL improvements, additional operators, etc.

See [ASSIGNMENT.md](ASSIGNMENT.md#6-rubric) for the full rubric.

## Resources

- [Lisp on Wikipedia](https://en.wikipedia.org/wiki/Lisp_(programming_language))
- [Go Type Assertions](https://go.dev/tour/methods/15)
- [Recursive Descent Parsing](https://en.wikipedia.org/wiki/Recursive_descent_parser)

## Getting Help

- Start by reading [ASSIGNMENT.md](ASSIGNMENT.md) thoroughly
- Implement in order: `parser.go` → `eval.go` → `main.go`
- Test frequently (after each small change)
- Use the skeleton code's TODO comments as a guide
- Check the expected output in `output/` directory

---

**Author**: Assignment design based on Philip Napoli and Nina Tacheva's mini interpreter project

**Due**: One week from assignment release
