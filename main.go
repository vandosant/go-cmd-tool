package main

import (
	"fmt"
	"os"
)

func init() {
	os.Args = append(os.Args, "-local", "u=admin", "--help")
}

func main() {
	// Program Name is the first argument
	cmd := os.Args[0]
	fmt.Printf("Program Name: %s\n", cmd)

	// Count the number of arguments
	argCount := len(os.Args[1:])
	fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)

	for i, a := range os.Args[1:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
	}
}
