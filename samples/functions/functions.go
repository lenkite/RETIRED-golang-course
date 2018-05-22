package main

import "fmt"

//function that takes two ints and returns their sum as an int. In
//this example, add takes two parameters of type int.  Notice that the
//type comes _after_ the variable name.
func add(x int, y int) int {
	return x + y
}

//For multiple consecutive parameters of the same type, the type name
//can be omitted like-typed parameters up to the final parameter that
//declares the type.
func add3(a, b, c int) int {
	return a + b + c
}

func add4(a, b, c int, d int8) int {
	return a + b + c + int(d)
}

//Go does NOT SUPPORT OVERLOADING

func main() {
	//Call a function just as you’d expect, with name(args).
	res := add(1, 2)
	fmt.Println("1+2 =", res)
	res = add3(1, 2, 3)
	fmt.Println("1+2+3 =", res)
	res = add4(1, 2, 3, -127)
	fmt.Println("1+2+3+-127=", res)
}
