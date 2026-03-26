.PHONY: all build clean test help

# Default target
all: build test

# Build the interpreter
build:
	cd src && go build -o ../interpreter .

# Run all tests and compare with expected output
test: build
	@echo "========================================="
	@echo "Running all tests..."
	@echo "========================================="
	@for testfile in input/*.lspgo; do \
		testname=$$(basename $$testfile .lspgo); \
		outfile="output/$${testname}.out"; \
		echo ""; \
		echo "Running test: $$testname"; \
		./interpreter "$$testfile" > "$${outfile}.tmp" 2>&1; \
		if diff -q "$${outfile}.tmp" "$$outfile" > /dev/null 2>&1; then \
			echo "✓ PASS: $$testname"; \
		else \
			echo "✗ FAIL: $$testname"; \
			echo "Expected:"; \
			cat "$$outfile"; \
			echo "Got:"; \
			cat "$${outfile}.tmp"; \
		fi; \
		rm -f "$${outfile}.tmp"; \
	done
	@echo ""
	@echo "========================================="
	@echo "Test run complete!"
	@echo "========================================="

# Clean up build artifacts
clean:
	rm -f interpreter
	find . -name "*.tmp" -delete

# Display help
help:
	@echo "Lisp Interpreter Assignment - Makefile targets"
	@echo ""
	@echo "make all     - Build and run tests (default)"
	@echo "make build   - Build the interpreter"
	@echo "make test    - Run all tests"
	@echo "make clean   - Clean build artifacts"
	@echo "make help    - Display this help message"
