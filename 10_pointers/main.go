package main

import "fmt"

func main() {
	originalNumber := 10
	pointerToOgNo := &originalNumber

	fmt.Println(pointerToOgNo)  // the memory address of "originalNumber"
	fmt.Println(*pointerToOgNo) // value inside it

	*pointerToOgNo = 15         // changing value of originalNumber by its pointer
	fmt.Println(originalNumber) // the value of originalNumber has been set to 15

}
