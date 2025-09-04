package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Cleaning a coffee machine...")
		fmt.Println("Suspending coffee machine...")
	}()

	fmt.Println("Brewing a fresh cup of espresso...")
	fmt.Println("Brewing a fresh cup of cappuccino...")
}
