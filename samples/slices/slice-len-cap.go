package main

import "fmt"

//The length of a slice is the number of elements it contains.
//The capacity of a slice is the number of elements in the
//underlying array, counting from the first element in the slice.
func main() {

	//The length and capacity of a slice s can be obtained using the
	//expressions len(s) and cap(s).

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("1> len=%d cap=%d %v\n", len(s), cap(s), s)

	// Slice the slice to give it zero length.
	s = s[:0]
	fmt.Printf("2> len=%d cap=%d %v\n", len(s), cap(s), s)

	// Extend its length.
	s = s[:4]
	fmt.Printf("3> len=%d cap=%d %v\n", len(s), cap(s), s)

	// Drop its first two values.
	s = s[2:]
	fmt.Printf("4> len=%d cap=%d %v\n", len(s), cap(s), s)
}
