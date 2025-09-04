package main

import "fmt"

func closeShop() {
	fmt.Println("Closing the coffee shop...")
}

func main() {
	defer closeShop()

	fmt.Println("Opening the coffee shop...")
	fmt.Println("Serving a customer...")
}
