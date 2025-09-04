package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Espresso":   2.50,
		"Latte":      3.75,
		"Cappuccino": 3.50,
		"Americano":  2.75,
	}

	fmt.Println("Menu today is:", menu)
	fmt.Printf("Latte costs: $%.2f\n", menu["Latte"])
	fmt.Printf("Americano costs: $%.2f\n", menu["Americano"])

	menu["Latte"] = 4.25 // change value of the key
	fmt.Printf("New Latte prcie : $%.2f\n", menu["Latte"])

	fmt.Println("Menu items quantity:", len(menu))

	menu["Mocha"] = 4.00 // add new key-value pair
	fmt.Println("Updated menu today is:", menu)
	fmt.Println("Updated menu items quantity:", len(menu))

}
