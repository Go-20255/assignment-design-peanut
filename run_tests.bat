@echo off
REM Run all .lspgo tests in the tests directory and print their outputs

for %%f in (tests\*.lspgo) do (
    echo Running %%f:
    pushd solution
    go run . "..\%%f"
    popd
    echo ---
)
