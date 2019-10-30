package main

import (
	"fmt"
	"math"

	"github.com/lets-go/03_packages/strutils"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))

	fmt.Println(strutils.Reverse("The quick brown fox jumped over the lazy dog."))
}
