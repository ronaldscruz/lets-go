package main

import "fmt"

func attack(attackerDamage, secondHeroHP int) int {
	return secondHeroHP - attackerDamage
}

func main() {
	attackerDamage := 15
	secondHeroHP := 40
	fmt.Printf("Dealt %d damage on hero2, his life is now: %d/%d", attackerDamage, attack(attackerDamage, secondHeroHP), secondHeroHP)
}
