package main

import "fmt"

func main() {
	// The zero value of a slice is nil.
	var s []int
	//A nil slice has a length and capacity of 0 and has no underlying array.
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
