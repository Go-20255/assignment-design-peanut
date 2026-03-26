# CSCI 541/641 Programming Skills: Go

# Homework Assignment: Build Your Own Lisp Interpreter

`Due date: One Week from Release`

<img src="./images/lisp-gopher.png" width="250">

---

## 1. Overview

Welcome to the **Great Parenthesis Crisis of 2026**! Your mission, should you choose to accept it (and you must, because grades), is to build a functional Lisp interpreter in Go. 

Lisp is one of the oldest and most elegant programming languages, famous for having exactly 1,000% more parentheses than any reasonable language needs, but also for teaching deep lessons about computation, evaluation, and the inevitable stack overflow errors.

In this assignment, you will build an interpreter that can:
- **Parse** Lisp code (lots of parentheses included)
- **Evaluate** expressions with proper semantics
- **Handle** variables, conditionals, functions, and lambdas
- **NOT crash** when your code is elegantly wrong

By the end of this week, you will understand why some programmers love Lisp, why others run screaming, and why there's a joke that "LISP = Lots of Irritating Superfluous Parentheses."

### Learning Objectives

* **Parsing & Tokenization**: Break source code into meaningful tokens and build abstract syntax trees
* **Type Systems**: Implement multiple types (numbers, booleans, symbols, lists) and operations on them
* **Environments & Scoping**: Manage variable bindings and function closures
* **Recursive Evaluation**: Implement an evaluator that recursively processes nested expressions
* **Functional Programming**: Understand first-class functions, lambdas, and higher-order programming
* **Error Handling**: Gracefully report and contextualize errors (or at least try to)

---

## 2. Requirements

### 2.1 Supported Language Features

Your Lisp interpreter must support the following:

#### Literals
- **Numbers**: Floating point numbers (e.g., `1`, `3.14`, `-42`)
- **Booleans**: `#t` (true) and `#f` (false)
- **Symbols**: Identifiers for variables and functions (e.g., `x`, `my_var`, `+`)

#### Operators
- **Arithmetic**: `+`, `-`, `*`, `/` (work on lists of operands)
  - Examples: `(+ 1 2 3)` → `6`, `(* 4 5)` → `20`
- **Comparison**: `eq?` (equality for all types)
  - Example: `(eq? 3 3)` → `#t`

#### Special Forms
- **`define`**: Bind a symbol to a value
  - Syntax: `(define symbol expr)`
  - Example: `(define x 7)` stores 7 in variable `x`
- **`if`**: Conditional evaluation
  - Syntax: `(if condition then-expr else-expr)`
  - Example: `(if #t 1 2)` → `1`
- **`lambda`**: Create anonymous functions
  - Syntax: `(lambda (params...) body)`
  - Example: `(lambda (x) (+ x 1))` creates an increment function

#### Advanced Features
- **Function Calls**: Call functions with arguments
  - Example: `(define inc (lambda (x) (+ x 1)))` then `(inc 5)` → `6`
- **Variable Lookup**: Resolve variables in the current environment
- **Closures**: Lambdas capture their defining environment

### 2.2 Program Input/Output

#### Input
Your interpreter accepts input via:
1. **Standard input mode** (REPL): Interactive prompt `> ` for line-by-line input
2. **File mode**: Pass a filename as a command-line argument

#### Output Format

For **file mode**, output each input followed by its result:
```
Input: (+ 1 2)
Output: 3

Input: (define x 7)
Output: [no output for define]

Input: x
Output: 7
```

For **REPL mode**, accept user input and display results:
```
Beginning REPL. Type "quit" to exit.
> (+ 1 2)
3
> (define x 7)
> x
7
> quit
```

#### Error Handling
If an error occurs, print a meaningful error message:
```
Parse error: failed to parse input.
Error: Encountered unmatched '('

Semantic error: failed to evaluate input.
Error: symbol was not found in dictionary
```

### 2.3 Performance & Correctness

All provided test cases must:
- **Complete successfully** without crashes
- **Produce correct output** matching expected results
- **Evaluate in reasonable time** (< 1 second for any test case)

---

## 3. Project Structure

```
assignment-design-peanut/
├── ASSIGNMENT.md           # This file
├── Makefile                # Build automation
├── run.sh                  # Test runner script
├── go.mod                  # Go module definition
├── go.sum                  # Go dependencies
│
├── src/                    # Student starter code
│   ├── main.go            # [TODO] Entry point - STUDENT MUST COMPLETE
│   ├── parser.go          # [TODO] Tokenization & parsing - STUDENT MUST COMPLETE
│   ├── eval.go            # [TODO] Expression evaluation - STUDENT MUST COMPLETE
│   ├── types.go           # Type definitions (provided, no changes needed)
│   ├── env.go             # Environment/builtins (provided, no changes needed)
│   └── go.mod
│
├── solution/              # Reference solution (for checking your work)
│   ├── main.go
│   ├── parser.go
│   ├── eval.go
│   ├── types.go
│   ├── env.go
│   └── go.mod
│
├── input/                 # Test input files
│   ├── 01_arithmetic.lspgo
│   ├── 02_subtraction.lspgo
│   ├── 03_multiplication.lspgo
│   ├── 04_division.lspgo
│   ├── 05_define.lspgo
│   ├── 06_if_true.lspgo
│   ├── 07_if_false.lspgo
│   ├── 08_eq.lspgo
│   ├── 09_lambda.lspgo
│   └── 10_lambda2.lspgo
│
└── output/                # Expected output files
    ├── 01_arithmetic.out
    ├── 02_subtraction.out
    ├── ... (one for each input)
    └── 10_lambda2.out
```

---

## 4. Implementation Guide

### 4.1 Core Components

Your interpreter consists of 5 main components:

#### 1. **Types** (`types.go` - PROVIDED)
Defines the data types used in the language:
```go
type Number float64
type Bool bool
type Symbol string
type Expr any
type List = []Expr
type Env = map[string]any
type Lambda struct {
    params []Symbol
    body   Expr
    env    Env
}
```

#### 2. **Environment** (`env.go` - PROVIDED)
Provides built-in functions like `+`, `-`, `*`, `/`, and `eq?`.

#### 3. **Parser** (`parser.go` - STUDENT TODO)
Must implement:
- **`tokenize(chars string) []string`**: Convert input string to tokens
  - Split on parentheses and whitespace
  - Example: `"(+ 1 2)"` → `[]string{"(", "+", "1", "2", ")"}`
- **`parse(chars string) (Expr, error)`**: Top-level parser
  - Validates complete parsing of all input
  - Returns error if extra tokens remain
- **`parse_tokens(tokens []string, idx *int) (Expr, error)`**: Recursive parser
  - **Skeleton hints**: 
    - Handle `(` by recursively parsing until `)`
    - Handle `)` as error (unexpected)
    - Handle default case: try parsing as number, boolean, or symbol

#### 4. **Evaluator** (`eval.go` - STUDENT TODO)
Must implement:
- **`eval(expr Expr, env Env) (any, error)`**: Core evaluation engine
  - **Skeleton hints**:
    - **Symbol**: Look up in environment
    - **Number/Bool**: Return as-is (self-evaluating)
    - **List**: 
      - Special forms: `define`, `if`, `lambda` (check first element)
      - Otherwise: Function call (evaluate first element, then arguments)
    - Use type assertions and switches

#### 5. **Main** (`main.go` - STUDENT TODO)
Must implement:
- **`repl(env Env)`**: Interactive prompt loop
  - Display `> ` prompt
  - Accept input, parse, evaluate, print result
  - Exit on "quit"
- **`file(filename string, env Env) error`**: Read and evaluate file
  - Parse and evaluate each top-level expression
  - Display input and output for each expression
- **`main()`**: Decide between REPL (no args) or file mode (with filename)

### 4.2 Implementation Steps (Recommended Order)

1. **Start with `types.go` and `env.go`** - These are provided; just understand them
2. **Implement `parser.go`**:
   - Start with `tokenize()` (easy, string manipulation)
   - Then `parse_tokens()` (recursive, uses pattern matching)
   - Then `parse()` (validation wrapper)
3. **Implement `eval.go`**:
   - Handle symbols, numbers, booleans (easy)
   - Handle arithmetic operators via function calls
   - Handle `define` special form
   - Handle `if` special form
   - Handle `lambda` special form (trickiest!)
4. **Implement `main.go`**:
   - REPL logic (reading, parsing, evaluating)
   - File reading and batch evaluation
   - Command-line argument handling

### 4.3 Key Implementation Tips

**Parsing:**
- Parentheses define list boundaries
- Whitespace separates tokens
- Numbers can be parsed with `strconv.ParseFloat()`
- Anything else is a symbol

**Evaluation - The Critical Insight:**
- All forms are consistent: `(form arg1 arg2 ...)`
- **Special forms** check the first element (e.g., `define`, `if`, `lambda`)
- **Regular functions** evaluate all arguments, then apply the function
- **Critical for lambdas**: When creating a lambda, capture the **current environment** so variables are resolved correctly when called later

**Environments:**
- Passing by reference (`Env` is a map) means modifications affect the caller
- When calling a lambda, create a **new environment** with the lambda's captured environment + parameter bindings
- When calling regular functions, evaluate arguments in the **current** environment

**Error Messages:**
- Be specific: "symbol was not found in dictionary" beats "ERROR"
- Include context: which symbol? which line?
- Users (including TAs grading) should understand what went wrong

---

## 5. Testing

### 5.1 Running Tests

```bash
# Build and run all tests
./run.sh

# Build only
make build

# Clean build artifacts
make clean
```

### 5.2 Test Cases

10 test cases are provided in `input/`, from simple arithmetic to advanced lambdas:

| Test | File | Concept | Example |
|------|------|---------|---------|
| 1 | `01_arithmetic.lspgo` | Basic `+` | `(+ 1 2)` → `3` |
| 2 | `02_subtraction.lspgo` | Basic `-` | `(- 10 3)` → `7` |
| 3 | `03_multiplication.lspgo` | Basic `*` | `(* 4 5)` → `20` |
| 4 | `04_division.lspgo` | Basic `/` | `(/ 20 4)` → `5` |
| 5 | `05_define.lspgo` | Variables | `(define x 7)` then `x` → `7` |
| 6 | `06_if_true.lspgo` | Conditional (true) | `(if #t 1 2)` → `1` |
| 7 | `07_if_false.lspgo` | Conditional (false) | `(if #f 1 2)` → `2` |
| 8 | `08_eq.lspgo` | Equality | `(eq? 3 3)` → `true` |
| 9 | `09_lambda.lspgo` | Simple lambda | Define and call `inc` function |
| 10 | `10_lambda2.lspgo` | Multi-param lambda | Define and call `add` function |

Expected output for each test is in `output/`.

### 5.3 Comparing Your Output

```bash
# After running tests, compare with expected:
diff output/01_arithmetic.out output_saved/01_arithmetic.out
```

---

## 6. Rubric

Total Points: **100**

### Code Quality & Implementation (60 points)

| Component | Points | Criteria |
|-----------|--------|----------|
| **Parser** | 15 | <ul><li>✓ Tokenization correctly splits input (3 pts)</li><li>✓ Handles parentheses and nesting (5 pts)</li><li>✓ Parses numbers, booleans, symbols correctly (7 pts)</li></ul> |
| **Evaluator - Basics** | 20 | <ul><li>✓ Symbol lookup works (4 pts)</li><li>✓ Numbers and booleans are self-evaluating (2 pts)</li><li>✓ Arithmetic operators `+`, `-`, `*`, `/` work (6 pts)</li><li>✓ Comparison operator `eq?` works (3 pts)</li><li>✓ Error handling for undefined symbols (5 pts)</li></ul> |
| **Evaluator - Special Forms** | 15 | <ul><li>✓ `define` creates variable bindings (5 pts)</li><li>✓ `if` evaluates conditionals correctly (5 pts)</li><li>✓ `lambda` creates function objects (5 pts)</li></ul> |
| **Evaluator - Advanced** | 10 | <ul><li>✓ Lambda functions can be called with arguments (5 pts)</li><li>✓ Closures capture the correct environment (5 pts)</li></ul> |

### Functional Requirements (25 points)

| Test Case | Points | Status |
|-----------|--------|--------|
| Tests 1-4 (Arithmetic) | 4 | ✓/✗ |
| Tests 5-7 (Variables & Conditionals) | 4 | ✓/✗ |
| Test 8 (Equality) | 4 | ✓/✗ |
| Tests 9-10 (Lambdas) | 13 | ✓/✗ |

### Code Style & Documentation (10 points)

| Criterion | Points |
|-----------|--------|
| Readable variable names & comments | 3 |
| Proper error messages | 3 |
| No unnecessary complexity | 2 |
| REPL mode works correctly | 2 |

### Bonus (up to 5 points)

- **+2**: REPL mode is fully functional and user-friendly
- **+1**: Handle edge cases gracefully (e.g., dividing by zero, empty lists)
- **+1**: Implement additional operators (e.g., `>`, `<`, `and`, `or`)
- **+1**: Add support for `list` and `car`/`cdr` operations

---

## 7. Common Mistakes to Avoid

1. **Forgetting to capture the environment in lambdas**: Lambdas must store `env` at creation time
2. **Modifying the global environment instead of a local copy**: When calling functions, use a new environment
3. **Not handling recursive expressions**: `(+ 1 (+ 2 3))` should work—evaluate arguments recursively
4. **Parser off-by-one errors**: Use pointers to `idx` carefully; increment **after** using the token
5. **Type assertion panics**: Use `ok` pattern, not bare assertions
6. **Forgetting to handle all cases**: Switch statements should cover all `expr.(type)` possibilities

---

## 8. Submission

1. Complete all **TODO** sections in `src/`
2. Ensure all 10 test cases pass: `./run.sh`
3. Verify REPL mode works
4. Commit and push to GitHub
5. All code must be in `src/` directory

**Deadline**: One week from assignment release

---

## 9. Resources

- [Lisp Wikipedia](https://en.wikipedia.org/wiki/Lisp_(programming_language))
- [Go Type Assertions](https://go.dev/tour/methods/15)
- [Go Interfaces](https://go.dev/tour/methods/9)
- [Recursive Descent Parsing](https://en.wikipedia.org/wiki/Recursive_descent_parser)

---

## 10. FAQ

**Q: Can I use reflection?**
A: You can, but it's overkill. Type assertions and switches are simpler.

**Q: How do I debug my parser?**
A: Print tokenized output, then recursively print the parsed AST structure.

**Q: Do I need to handle all edge cases?**
A: No, but the 10 provided tests should all pass.

**Q: Will the solution work with the Go Playground?**
A: Yes, REPL mode will. File mode requires `os` package, which isn't available on playground.

**Q: Is this your first time writing an interpreter?**
A: That's okay! You've got this. Start small, test frequently, and celebrate when test 1 passes. 🎉

