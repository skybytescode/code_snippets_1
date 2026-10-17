package main

import (
	"fmt"
	"os"
)

// Usage: go run . [name]
func main() {
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Printf("Hello, %s!\n", name)
}
