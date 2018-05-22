// samples/structs/structs.go
package main

import "fmt"

//This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

func main() {
	//Struct Literal
	fmt.Println(person{"Bob", 20})

	//You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	//	Omitted fields will be zero-valued.
	fmt.Println(person{name: "Fred"})

	//An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	//Access struct fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	//You can also use dots with struct pointers - the pointers are automatically
	//dereferenced.
	sp := &s //you could declare tihs like var sp *person
	fmt.Println(sp.age)
	fmt.Printf("Type of sp: %T\n", sp)

	//Structs are mutable
	(*sp).age = 42
	fmt.Println(sp.age)
	//For convenience. You have implicit pointer de-reference. No need for C
	// arrow -> operator
	sp.age = 99
	fmt.Println(sp.age)
}
