// samples/panic/panic.go

package main

import "os"

func main() {
	// A common use of panic is to abort if a function returns an error value
	//that we don’t know how to (or want to) handle. Here’s an example of panicking
	//if we get an unexpected error when creating a new file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
