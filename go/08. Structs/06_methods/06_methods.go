package main

import "fmt"

type CoffeeShop struct {
	Name string
}

// method with value receiver
func (shop CoffeeShop) Greet() {
	fmt.Println("Welcome to the", shop.Name)
}

func main() {
	myShop := CoffeeShop{
		Name: "Brew & Beans",
	}
	myShop.Greet()

}
