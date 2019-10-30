package main

import "fmt"

func main() {
	emailList := make(map[string]string)

	emailList["Ronald"] = "ronald.cruz@grupokrs.com.br"
	emailList["Erico Rocha"] = "erico.rocha@mkt.com"
	emailList["Carl Johnson"] = "cj_grove4life@gtasa.com"

	// Shorthand syntax
	heroesAndPower := map[string]int{"Superman": 10, "Batman": 7, "Hulk": 12}
	fmt.Println(heroesAndPower)

	fmt.Println(emailList)

	nameFromQuery := "Erico Rocha" // erico.rocha@mkt.com
	fmt.Println(emailList[nameFromQuery])

	delete(emailList, nameFromQuery)
	fmt.Printf("%s deleted successfully!\n", nameFromQuery)
}
