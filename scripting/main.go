//usr/local/bin/go run "$0" "$@"
package main

import (
	"fmt"
	"os"
)

func main() {
	s := "World"

	// beacuse default argument size is one, os.Args[0] is script name
	if len(os.Args) == 1 {
		fmt.Println("Pls use one argument")
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		s = os.Args[1]
	}
	fmt.Printf("Hello, %v\n", s)
	if s == "fail" {
		os.Exit(30)
	}
}


