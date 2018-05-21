package main

import (
	"fmt"
	"math/cmplx"
)

var ( //var declarations may be factored into blocks like import
	ToBe     bool       = false
	Max16Int uint16     = 1<<16 - 1
	Max32Int uint32     = 1<<32 - 1
	MaxInt   uint64     = 1<<64 - 1
	z        complex128 = cmplx.Sqrt(-5 + 12i)
	//Ideally should omit type in above as can be inferred
)

func main() {
	fmt.Printf("(ToBe) Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("(Max16Int) Type: %T Value: %v\n", Max16Int, Max16Int)
	fmt.Printf("(Max32Int) Type: %T Value: %v\n", Max32Int, Max32Int)
	fmt.Printf("(MaxInt) Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("(z) Type: %T Value: %v\n", z, z)
}
