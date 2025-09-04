package main

import "fmt"

func main() {
	menu := map[string]float64{
		"Latte":      3.75,
		"Espresso":   2.50,
		"Cappuccino": 3.50,
		"Americano":  2.75,
	}
	fmt.Println("Original menu:", menu)

	delete(menu, "Latte") // remove key with value
	fmt.Println("Updated menu:", menu)

	delete(menu, "Americano") // remove key with value
	fmt.Println("Updated menu:", menu)

	delete(menu, "Tea") // if such key is absent nothing happens, no error
}
