#!/bin/bash
#
# Lisp Interpreter Assignment - Test Runner Script
# This script builds the interpreter and runs all tests
#

set -e  # Exit on error

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Counters
PASSED=0
FAILED=0

echo "========================================="
echo "Lisp Interpreter - Test Suite"
echo "========================================="
echo ""

# Build the interpreter
echo "Building interpreter..."
cd src
go build -o ../interpreter .
cd ..

if [ ! -f interpreter ]; then
    echo -e "${RED}✗ Build failed!${NC}"
    exit 1
fi
echo -e "${GREEN}✓ Build successful${NC}"
echo ""

# Run tests
echo "========================================="
echo "Running tests..."
echo "========================================="
echo ""

for testfile in input/*.lspgo; do
    testname=$(basename "$testfile" .lspgo)
    outfile="output/${testname}.out"
    tmpfile="${outfile}.tmp"
    
    # Run the test
    if ./interpreter "$testfile" > "$tmpfile" 2>&1; then
        # Compare with expected output
        if diff -q "$tmpfile" "$outfile" > /dev/null 2>&1; then
            echo -e "${GREEN}✓ PASS${NC}: $testname"
            ((PASSED++))
        else
            echo -e "${RED}✗ FAIL${NC}: $testname"
            echo "  Expected:"
            sed 's/^/    /' "$outfile"
            echo "  Got:"
            sed 's/^/    /' "$tmpfile"
            ((FAILED++))
        fi
    else
        echo -e "${RED}✗ ERROR${NC}: $testname (failed to run)"
        cat "$tmpfile"
        ((FAILED++))
    fi
    rm -f "$tmpfile"
done

echo ""
echo "========================================="
echo -e "Results: ${GREEN}$PASSED passed${NC}, ${RED}$FAILED failed${NC}"
echo "========================================="

# Exit with error code if any tests failed
[ $FAILED -eq 0 ]
