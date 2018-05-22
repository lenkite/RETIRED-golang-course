package main

import "fmt"

// Slices can be created with the built-in make function
// The make function allocates a zeroed array and returns a slice that
// refers to that array

func main() {
	a := make([]int, 5)
	fmt.Printf("a. len=%d cap=%d %v\n", len(a), cap(a), a)

	b := make([]int, 0, 7)
	fmt.Printf("b. len=%d cap=%d %v\n", len(b), cap(b), b)

	c := b[:2]
	fmt.Printf("c. len=%d cap=%d %v\n", len(c), cap(c), c)

	d := c[2:5]
	fmt.Printf("d. len=%d cap=%d %v\n", len(d), cap(d), d)

	//Slices can be composed into multi-dimensional data structures. The
	//length of the inner slices can vary, unlike with multi-dimensional
	//arrays.
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
