// samples/for/for.go
package main

import "fmt"

func main() {
	//No parenthesis needed for go for loops

	//The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Printf("i=%d\n", i)
		i = i + 1
	}
	//A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Printf("j=%d\n", j)
	}
	//for without a condition will loop repeatedly until you break out
	//of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}
	//You can also continue to the next iteration of the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Printf("n=%d\n", n)
	}
	fmt.Println("For outer/inner scope")
	//for implicity defines *2* scopes
	for i, z := 0, 0; z < 5; z++ { //outerscope before {
		i = z
		z := 2 //inner scope
		fmt.Printf("Iteration: #%d z is %d\n", i, z)
	}
	fmt.Println("LOLWUT How did I exit?")
}
