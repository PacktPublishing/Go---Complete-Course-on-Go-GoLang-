package main

import "fmt"

func getDrinkInfo(customerName string, drink string, price float64) {
	info := "%s's favorite drink is %s and it's price is $%.2f\n"
	fmt.Printf(info, customerName, drink, price)
}

func main() {
	getDrinkInfo("Bogdan", "Capuccino", 4.50)
	getDrinkInfo("Alice", "Espresso", 2.50)
}
