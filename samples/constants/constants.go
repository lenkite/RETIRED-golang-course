package main

import (
	"fmt"
	"math"
)

//const declares a constant value that can be 'high-precision'
const s string = "constant"

const ( //const declarations can be facgtored
	// Big is a huge number which is 1 bit left-shifted 100 places.
	// the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Small is a small number which is Big shifted right 99 places,
	// so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	// A number can be given a type by using it in a context that
	// requires one, such as a variable assignment or function call.
	// For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(Small))
	fmt.Println(math.Sin(Big))
}
