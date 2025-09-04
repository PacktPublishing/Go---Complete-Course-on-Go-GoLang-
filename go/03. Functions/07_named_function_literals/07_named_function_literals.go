package main

import "fmt"

func main() {
	taxRate := 0.10 // 10% tax

	calculateTax := func(amount float64) float64 {
		return amount * taxRate
	}

	subtotal := 25.00
	tax := calculateTax(subtotal)
	total := subtotal + tax
	fmt.Printf("Total amount to pay: $%.2f\n", total)
}
