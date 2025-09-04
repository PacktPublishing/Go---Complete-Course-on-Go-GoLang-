package main

import "fmt"

func main() {
	coffeeTypes := [3]string{"Espresso", "Latte", "Cappuccino"}
	fmt.Println("Types of coffee:", coffeeTypes)
	fmt.Println("Length of the array:", len(coffeeTypes))

	coffeeTypes[len(coffeeTypes)-1] = "Milk" // access last element in the array
	fmt.Println("Types of coffee:", coffeeTypes)

	fmt.Println("String length is:", len("This is a coffee string!"))
}
