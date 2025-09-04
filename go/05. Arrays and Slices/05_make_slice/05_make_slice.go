package main

import "fmt"

func main() {
	ratings := []int{5, 4, 5, 5, 3}
	fmt.Println("Original ratings", ratings)

	ratings[2] = 3
	fmt.Println("Changed element with index 2:", ratings)

	fmt.Println("Length of the slice is:", len(ratings))

	coffeeTypes := make([]string, 3)
	fmt.Println("Coffee types after make:", coffeeTypes)

	coffeeTypes[0] = "Cappuccino"
	coffeeTypes[1] = "Latte"
	coffeeTypes[2] = "Espresso"

	fmt.Println("Coffee types after reassignment:", coffeeTypes)

}
