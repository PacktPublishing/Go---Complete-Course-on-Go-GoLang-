package main

import "fmt"

type CoffeeShop struct {
	Name  string
	Greet func(shop CoffeeShop)
}

func greetShop(shop CoffeeShop) {
	fmt.Println("Welcome to the", shop.Name)
}

func main() {
	myShop := CoffeeShop{
		Name:  "Brew & Beans",
		Greet: greetShop,
	}

	myShop.Greet(myShop)
}
