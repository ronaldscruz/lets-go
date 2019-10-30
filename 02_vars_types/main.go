package main

import "fmt"

func main() {
	// strings and integers doesnt need type annotations
	var name = "Pscythe"
	var age = 18

	// shorthand syntax to create variables
	weight := 72.6

	fmt.Println(name, age, weight)
	fmt.Printf("%T\n", name)   // string
	fmt.Printf("%T\n", age)    // int
	fmt.Printf("%T\n", weight) // float64

	// consts cant be changed
	const pi complex64 = 3.141592653589793238462643383279502884197169399375105820974944592307816406286
	fmt.Println(pi)

}
