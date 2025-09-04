package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso":   2.50,
		"Latte":      3.75,
		"Cappuccino": 3.50,
		"Americano":  2.75,
	}

	for drink, price := range menu {
		fmt.Printf("%s costs $%.2f\n", drink, price)
	}
}
