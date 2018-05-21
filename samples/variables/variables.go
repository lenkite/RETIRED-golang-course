package main

import "fmt"

func main() {
	// var declares 1 or more variables.
	var a = "initial"
	fmt.Println(a)

	// You can declare multiple variables at once.
	var b, c int = 1, 2 //type is declared last
	fmt.Println(b, c)

	// Go infers type of initialized variables from literal.
	var d = true //infer type as bool
	fmt.Println(d)

	// Inference also works for multiple variable declaration
	var clang, python, java = true, false, "no!"
	fmt.Println(clang, python, java)

	// Variables declared without a corresponding initialization
	// are ZERO-VALUED.
	// For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	var p int
	var q float64
	var r bool
	var s string
	fmt.Printf("(Zero values) p=%v q=%v r=%v s=%q\n", p, q, r, s)
	//zero value of s is actually emptoy string without quote

	// := syntax is shorthand for declaring and initializing
	f := "short" //equivalent to var f string = "short"
	fmt.Println(f)
	// := valid only INSIDE a function, not at package level

}
