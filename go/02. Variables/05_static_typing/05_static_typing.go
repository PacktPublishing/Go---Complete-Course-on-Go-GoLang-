package main

import "fmt"

func main() {
	// Explicit type declaration
	var cupsQty int = 3

	// cupsQty = "four" // ❌ Compile-time error
	fmt.Println("Number of cups:", cupsQty)

	// Implicit type declaration
	var wasProcessed = true

	// wasProcessed = "no" // ❌ Compile-time error
	fmt.Println("Order was processed:", wasProcessed)
}
