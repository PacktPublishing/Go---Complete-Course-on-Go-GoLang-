package main

import "fmt"

func DispenseCoffee(coffeeAmount int, cups int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Machine error:", r)
		}
	}()

	fmt.Printf("Dispensing %d grams of coffee into %d cups...\n", coffeeAmount, cups)
	amountPerCup := coffeeAmount / cups
	fmt.Printf("Each cup gets %d grams of coffee\n", amountPerCup)
}

func main() {
	fmt.Println("Starting coffee machine...")
	DispenseCoffee(750, 200)

	fmt.Printf("\nCoffee machine is still running...\n")
	DispenseCoffee(340, 0) // error is handled using recover()

	fmt.Printf("\nCoffee machine is still running...\n")
	DispenseCoffee(500, 150)
}
