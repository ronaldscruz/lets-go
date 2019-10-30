package main

import "fmt"

// Person : creates an generic and duplicable person
type Person struct {
	name string
	age  int
}

func main() {
	peopleInRoom := []Person{{"Ronald", 18}, {"Deschamps", 23}, {"Xulio", 22}, {"Petinho", 25}, {"Pruno", 24}}

	fmt.Println(peopleInRoom)
}
