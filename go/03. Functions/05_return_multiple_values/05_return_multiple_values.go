package main

import "fmt"

func processPayment(orderTotal float64, tip float64, amountPaid float64) (float64, float64) {
	totalAmountDue := orderTotal + tip
	change := amountPaid - totalAmountDue
	return totalAmountDue, change
}

func main() {
	// 6.50 (orderTotal) + 2.00 (tip) = 8.50 (totalAmountDue)
	// 10.00 (amountPaid)
	// 10.00 (amountPaid) - 8.50 (totalAmountDue) = 1.50 (change)

	// We need to calc totalAmountDue and change
	totalCost, change := processPayment(6.50, 2.00, 10.00)
	fmt.Printf("Total cost (with tip): $%.2f\n", totalCost)
	fmt.Printf("Change returned to the customer: $%.2f\n", change)

	fmt.Println("____________________")

	totalCost, change = processPayment(28.50, 1.50, 50.00)
	fmt.Printf("Total cost (with tip): $%.2f\n", totalCost)
	fmt.Printf("Change returned to the customer: $%.2f\n", change)
}
