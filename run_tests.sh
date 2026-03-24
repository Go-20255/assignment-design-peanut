#!/bin/bash
# Run all .lspgo tests in the tests directory and print their outputs

set -e
cd "$(dirname "$0")"

for testfile in tests/*.lspgo; do
    echo "Running $testfile:"
    go run . "$testfile"
    echo "---"
done
