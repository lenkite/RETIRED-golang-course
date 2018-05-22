// samples/cli1/cliargs.go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Program Name is always the first (implicit) argument
	cmd := os.Args[0]
	fmt.Printf("Program Name: %s\n", cmd)

	argCount := len(os.Args[1:])
	fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)

	fmt.Println("Enumeratig Arguments")
	for i, a := range os.Args[1:] {
		fmt.Printf("Argument %d is %s\n", i+1, a)
	}
}
