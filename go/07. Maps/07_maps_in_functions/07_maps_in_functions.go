package main

import "fmt"

func sellCoffee(stock map[string]int, coffeeType string, quantity int) {
	currentStockQty, exists := stock[coffeeType]

	if exists {
		if currentStockQty >= quantity {
			stock[coffeeType] = stock[coffeeType] - quantity
			fmt.Printf("Sold %d %s(s)\n", quantity, coffeeType)

			fmt.Println("\nsellCoffee function:")
			fmt.Println("    stock:", stock)
			fmt.Printf("    Location in memory %p\n\n", &stock) // location of the variable, not map itself!

		} else {
			fmt.Printf("Not enough %s in stock. Only %d left. %d ordered\n", coffeeType, currentStockQty, quantity)
		}

	} else {
		fmt.Printf("%s is not available in stock.\n", coffeeType)
	}
}

func main() {
	coffeeStock := map[string]int{
		"Espresso":   10,
		"Latte":      5,
		"Cappuccino": 8,
	}

	sellCoffee(coffeeStock, "Espresso", 2)
	sellCoffee(coffeeStock, "Cappuccino", 4)
	sellCoffee(coffeeStock, "Mocha", 1) // Not available
	sellCoffee(coffeeStock, "Latte", 7) // Not enough in stock

	fmt.Println("\nmain function:")
	fmt.Println("    coffeeStock:", coffeeStock)
	fmt.Printf("    Location in memory %p\n\n", &coffeeStock) // location of the variable, not map itself!
}
