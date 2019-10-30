package main

import "fmt"

func main() {
	// the only loop in Go is "for"
	for i := 10; i > 0; i-- {
		fmt.Printf("%d!\n", i)
	}

	fmt.Println("--- GOOO! ---")
}
