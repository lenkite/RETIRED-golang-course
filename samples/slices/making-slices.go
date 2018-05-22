// samples/slices/making-slices.go
package main

import "fmt"

// Slices can be created with the built-in make function
// The make function allocates a zeroed array and returns a slice that
// refers to that array

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 7)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

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

func printSlice(n string, s []int) {
	fmt.Printf("%q: len=%d cap=%d %v\n", n, len(s), cap(s), s)
}
