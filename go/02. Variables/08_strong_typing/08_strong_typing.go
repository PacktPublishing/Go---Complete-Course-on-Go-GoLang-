package main

import "fmt"

func main() {
	// price of one cup of coffee
	price := 4.50 // float64

	// cups sold in one day
	quantity := 15 // int

	// total income
	// total := price * quantity // ❌ Go doesn't convert types automatically
	total := price * float64(quantity)

	fmt.Printf("Total income during a day %.2f\n", total)
}
