@echo off
REM Run all .lspgo tests in the tests directory and print their outputs

setlocal enabledelayedexpansion
set testnum=0

for %%f in (tests\*.lspgo) do (
    set /a testnum+=1
    echo TEST !testnum!
    pushd solution
    go run . "..\%%f"
    popd
    echo.
)
