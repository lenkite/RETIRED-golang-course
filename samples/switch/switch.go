// samples/switch/switch.go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// Here’s a basic switch.
	i := 2
	fmt.Print("> Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//You can use commas to separate multiple expressions in the same
	//case statement. We use the optional default case in this example
	//as well.

	fmt.Println("> Weekday or Weekend ?")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// case can contain expressions matched against switch expression
	fmt.Println("> When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//switch without an expression is an alternate way to express if/else logic.
	//Here we also show how the case expressions can be non-constants.
	//clean way to write long if-then-else chains

	fmt.Println("> Before or After noon?")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	fmt.Print("> Go running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	//TODO: cover type switch later
	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm a bool")
	// 	case int:
	// 		fmt.Println("I'm an int")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }
	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")

}
