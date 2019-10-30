package main

import "fmt"

// Person : creates an generic and duplicable person
type Person struct {
	name string
	age  int
}

func main() {
	peopleInRoom := []Person{{"Ronald", 18}, {"Xiofane", 27}, {"Xulio", 22}, {"Petinho", 25}, {"Pruno", 24}}

	for personIndex, person := range peopleInRoom {
		if person.name == "Petinho" {
			fmt.Printf("PETINHO NO LUGAR %d\n", personIndex)
		} else {
			fmt.Printf("%s :l\n", person.name)
		}
	}

	// if you don't want to get the index, just place an underscore on its name
	for _, person := range peopleInRoom {
		fmt.Printf("> %s\n", person.name)
	}

	// you can also iterate through maps by passing "key, value" as left-side parameters of range
	contactList := map[string]string{"Ronald": "11 94838-9293", "Bruna": "21 98830-3019", "Toninho": "12 99289-0990"}

	for name, phoneNo := range contactList {
		fmt.Printf("Name: %s\nNumber: %s\n\n", name, phoneNo)
	}
}
