package main

import "fmt"

func addSheepCount() func() int {
	currentSheepCount := 0

	return func() int {
		currentSheepCount++
		return currentSheepCount
	}
}

func main() {

	startCountingSheeps := addSheepCount()

	fmt.Println(startCountingSheeps()) // 1
	fmt.Println(startCountingSheeps()) // 2
	fmt.Println(startCountingSheeps()) // 3
	fmt.Println(startCountingSheeps()) // 4

	// resets the closure (currentSheepCount := 0)
	restartSheepCount := addSheepCount()

	fmt.Println(restartSheepCount()) // 1

}
