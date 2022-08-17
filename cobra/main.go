package main

import "github.com/russellcxl/go-algo/cobra/cmd/stringer"

/*

pkg - contains logic for functions
cmd - will create cobra commands and tag an alias to a function in pkg

to see all commands:
	go run ginkgo_test.go --help

to reverse a string:
	go run ginkgo_test.go reverse "hello"
	go run ginkgo_test.go rev "hello"

to inspect a string:
	go run ginkgo_test.go inspect "abc123"
	go run ginkgo_test.go inspect "abc123" --digits

*/

func main() {
	stringer.Execute()
}
