package main

import (
	"fmt"
	"os"
)

func main() {
	// Program Name is the first argument
	cmd := os.Args[0]
	fmt.Printf("Program Name: %s\n", cmd)

	// Count the number of arguments
	argCount := len(os.Args[1:])
	fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)
}
