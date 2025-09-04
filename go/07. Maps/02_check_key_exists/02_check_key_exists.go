package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso": 2.50,
	}

	drink := "Cappuccino"
	fmt.Printf("%s price: $%.2f\n", drink, menu[drink])

	price, exists := menu[drink]

	fmt.Println("Exists:", exists)

	if exists {
		fmt.Printf("%s costs $%.2f\n", drink, price)
	} else {
		fmt.Printf("%s is not on the menu\n", drink)
	}

}
