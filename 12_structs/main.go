package main

import (
	"fmt"
	"strconv"
)

// Customer : a customer base
type Customer struct {
	fullName, city, gender string
	age                    int
	balance                float64
}

func (c Customer) printableCustomerData() string {
	return "Name: " + c.fullName + " | City: " + c.city + " | Gender: " + c.gender + " | Age: " + strconv.Itoa(c.age)
}

func (c *Customer) boughtProduct(price float64) {
	c.balance += price
}

func main() {
	customerGerald := Customer{"Gerald Remotsuc", "New York", "M", 32, 0.0}
	fmt.Println(customerGerald.printableCustomerData())

	customerGerald.boughtProduct(20)
	fmt.Println(customerGerald.balance)
}
