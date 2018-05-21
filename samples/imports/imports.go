package main

import ( // "factorized" import statement for multiple imports
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//function must be prefixed by pkg name. No static imports
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
