package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Brew&Beans Terminal!")

	for { // infinite loop
		var order string
		fmt.Print("Enter your coffee name (or type 'exit' to quit): ")
		fmt.Scanln(&order)

		if order == "exit" {
			fmt.Println("Thank you for visiting Brew&Beans. Good bye!")
			break
		}

		if order == "" {
			fmt.Println("Please enter a valid order")
			continue
		}

		fmt.Println("Preparing your order...", order)
	}

	fmt.Println("Finishing program...")

}
