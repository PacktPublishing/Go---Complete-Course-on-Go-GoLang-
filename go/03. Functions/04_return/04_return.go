package main

import "fmt"

func updateTotalPoints(currentPoints int, newPoints int) int {
	return currentPoints + newPoints
}

func calculateLoyaltyPoints(amountSpent float64) int {
	loyaltyPoints := int(amountSpent * 2)
	return loyaltyPoints
}

func main() {
	totalPoints := 120
	var newlyEarnedPoints int = calculateLoyaltyPoints(9.30) // 18
	fmt.Println("Earned points today:", newlyEarnedPoints)

	totalPoints = updateTotalPoints(totalPoints, newlyEarnedPoints)

	fmt.Println("Updated loyalty points:", totalPoints) // 120 + 18 = 138
}
