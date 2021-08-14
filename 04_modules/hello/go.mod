module example.com/hello

go 1.16

// We want to use the local module, for that: go mod edit -replace example.com/greetings=../greetings
// After this run: go mod tidy
// In production we will remove this replace
replace example.com/greetings => ../greetings

// In production we will remove this require
require example.com/greetings v0.0.0-00010101000000-000000000000
