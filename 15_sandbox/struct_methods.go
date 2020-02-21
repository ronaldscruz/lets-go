package main

import (
	"fmt"
)

// Hero : -
type Hero struct {
	name        string
	health      float64
	attackPower float64
}

// Attack : -
func (h *Hero) Attack() float64 {
	fmt.Printf("Hero %s attacked and dealt %.1f damage! \nHe got empowered with +1 attack dmg \n\n", h.name, h.attackPower)

	h.attackPower++

	return h.attackPower
}

func main() {
	hero1 := Hero{
		name:        "Ronald",
		health:      100,
		attackPower: 22,
	}

	hero1.Attack()
	hero1.Attack()
	hero1.Attack()
	hero1.Attack()
}
