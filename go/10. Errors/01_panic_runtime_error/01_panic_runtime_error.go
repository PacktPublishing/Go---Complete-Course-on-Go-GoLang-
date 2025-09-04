package main

import "fmt"

func DispenseCoffee(coffeeAmount int, cups int) {
	fmt.Printf("Dispensing %d grams of coffee into %d cups...\n", coffeeAmount, cups)
	amountPerCup := coffeeAmount / cups
	fmt.Printf("Each cup gets %d grams of coffee\n", amountPerCup)
}

func main() {
	fmt.Println("Starting coffee machine...")

	DispenseCoffee(750, 200)

	fmt.Println("Coffee machine is still running...")

	DispenseCoffee(340, 0) // panic: runtime error: integer divide by zero
}
