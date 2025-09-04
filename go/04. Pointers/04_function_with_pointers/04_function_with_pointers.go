package main

import "fmt"

func applyDiscount(price *float64, discountRate float64) {
	*price = *price - (*price * discountRate)
}

func main() {
	var coffeePrice float64 = 5.00
	var discount float64 = 0.10
	fmt.Printf("Basic coffee price: $%.2f\n", coffeePrice) // 5.00

	applyDiscount(&coffeePrice, discount)
	fmt.Printf("Price with discount: $%.2f\n", coffeePrice) // 4.50
	// 5.00
	// 10%
	// 5.00 - 5.00 * 0.10 = 5.00 - 0.50 = 4.50
}
