// samples/cli2/flags.go

package main

import (
	"flag"
	"fmt"
)

func main() {
	var port int
	// flag that is of type int with a default value of 8000, and usage text
	flag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000.")
	flag.Parse()

	fmt.Printf("port = %d", port)
	fmt.Printf("other args: %+v\n", flag.Args())
}
