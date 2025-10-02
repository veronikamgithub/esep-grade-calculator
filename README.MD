# ESEP Grade Calculator

This library is a simple utility to calculate your grades.


There are 3 included files:
    - go.mod specifies the module and version number for golang
    - grade_calculator.go provides the functions and code, along with the structs for grade calculation and       variables
    - grade_calculator_test.go contains some tests with expected values where certain values will be passed in, ran against the functions, and a line which causes an error to be thrown if the unexpected value doesn't equate to the calculated one. 
    - make sure to install go from the go.dev and also add it to the windows environment variables on the computer. 



Runtime requirements:
Golang 1.24+

Run unit tests:
```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:
```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```