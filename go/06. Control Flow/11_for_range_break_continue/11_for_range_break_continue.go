package main

import "fmt"

func main() {
	drinks := []string{"Latte", "Expired Milk", "Espresso", "Matcha", "Cappuccino"}

	for _, drink := range drinks {
		if drink == "Expired Milk" {
			fmt.Println("Warning! Skipping expired milk")
			continue
		}

		if drink == "Matcha" {
			fmt.Println("Matcha sold out! Stopping service...")
			break
		}

		fmt.Println("Preparing drink:", drink)
	}

}
