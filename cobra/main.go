package main

import "github.com/russellcxl/go-algo/cobra/cmd/stringer"

/*

pkg - contains logic for functions
cmd - will create cobra commands and tag an alias to a function in pkg

to see all commands:
	go run main.go --help

to reverse a string:
	go run main.go reverse "hello"
	go run main.go rev "hello"

to inspect a string:
	go run main.go inspect "abc123"
	go run main.go inspect "abc123" --digits

*/

func main() {
	stringer.Execute()
}
