# Assignment Design Activity - Mini Interpreter

## Overview
Our programming assignment is a mini interpreter for a small Lisp-like language.

## Authors
This assignment was designed and created by Philip Napoli and Nina Tacheva.

## Assignment Details
Primary assignment details can be found in the ASSIGNMENT.md file.

As a short summary, this project asks students to build out portions of a Lisp interpreter from scratch. The final Lisp implementation supports:
- Numbers
- Symbols
- Booleans
- Definitions
- Branching with if
- Lambdas

Several unit tests are provided as well for students to test their implementation against. These are located in the tests/ directory.

The primary components students must implement are the parser (parser.go), evaluator (eval.go), and REPL and file executor (main.go).

## Issues/Limitations
Our environment does not have a pointer to the parent, so at the moment, nested scope is not fully supported. This is one thing we would ideally implement if we wanted to support recursion.