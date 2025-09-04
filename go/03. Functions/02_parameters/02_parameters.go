package main

import "fmt"

func greet(coffeeShopName string) {
	var greetMessage string = "Welcome to the Coffee Shop"
	fmt.Println(greetMessage, coffeeShopName)
}

func main() {
	greet("Brew & Beans")
	greet("Coffee & Milk")
}
