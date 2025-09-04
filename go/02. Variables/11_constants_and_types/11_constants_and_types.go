package main

import "fmt"

func main() {
	// Untyped constant with integer value
	const rewardPoints = 10

	fmt.Printf("Default type of rewardPoints is %T\n", rewardPoints) // int

	var totalRewardPoints float64 = 150.3

	// Adding untyped constant to a float64 - valid, constant adapts
	totalRewardPoints = totalRewardPoints + rewardPoints

	fmt.Printf("Updated loyalty points %.2f\n", totalRewardPoints)
}
