package main

import (
	"fmt"
	"os"
)

func init() {
	os.Args = append(os.Args, "-local", "u=admin", "--help")
}

func main() {
	cmd := len(os.Args[0])
	argCount := len(os.Args[1:])

	fmt.Printf("Program Name: %s\n", cmd)
	fmt.Printf("Total Arguments: %d\n", argCount)

	for i, a := range os.Args[1:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
	}
}
