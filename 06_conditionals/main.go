package main

import "fmt"

// Pokemon : creates an generic pokemon
type Pokemon struct {
	name  string
	power int
}

func main() {
	pikachu := Pokemon{"Pikachu", 10}
	vulpix := Pokemon{"Vulpix", 12}

	// Go conditionals doesn't require parenthesis
	if pikachu.power > vulpix.power {
		fmt.Println("Pikachu is better than Vulpix")
	} else {
		fmt.Println("Vulpix is better than Pikachu")
	}

	// Switch statements have an automatic "break" after each "case"
	customerPlan := "VIP Simple"

	switch customerPlan {
	case "VIP Tester":
		fmt.Println("Your plan is VIP Tester")

	case "VIP Simple":
		fmt.Println("Your plan is VIP Simple")

	case "VIP Plus":
		fmt.Println("Your plan is VIP Plus")

	default:
		fmt.Println("Invalid plan")
	}
}
