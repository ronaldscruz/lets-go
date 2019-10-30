package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Watermelon"
	fruitArr[1] = "Tomato"

	// Shorthand syntax
	carsArr := [2]string{"Voyage", "Fusca"}
	fmt.Println(carsArr)

	// "Unlimited" array
	skillsArr := []string{"Poison Bomb", "Water Gun", "Meteor", "Tornado"}
	fmt.Printf("\nThere are %d skills registered", len(skillsArr))
	fmt.Printf("\nThe last 2 are: %s\n", skillsArr[2:4])
}
