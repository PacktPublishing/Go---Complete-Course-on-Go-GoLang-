package main

import "fmt"

func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffee price", coffeePrice)
	// STEP 1
	// Compile Time (code you write):    var coffeePrice = 4.50
	// Runtime (what machine sees):      [some memory address] = 4.50

	// STEP 2
	// Compile Time (code you write):    fmt.Println("Coffee price", coffeePrice)
	// Runtime (what machine sees):      fmt.Println([some mem address], [memory address (same as on step 1)])
	fmt.Println("Memory address of price 4.50", &coffeePrice)

	coffeePrice = 5.00
	fmt.Println("Memory address of price 5.00", &coffeePrice)

	// pointerToCoffeePrice := &coffePrice  // same as next line
	var pointerToCoffeePrice *float64 = &coffeePrice

	/* go to the memory location where pointerToCoffeePrice points to
	and change value in this memory location */
	*pointerToCoffeePrice = 7.50

	fmt.Println("Updated coffeePrice value in memory", coffeePrice)
}
