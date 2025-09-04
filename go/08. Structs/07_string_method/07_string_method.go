package main

import "fmt"

type CoffeeType string

func (coffee CoffeeType) Describe() {
	fmt.Println("This is delicios", coffee)
}

func main() {
	var myCoffee CoffeeType = "Espresso"
	myCoffee.Describe()

}
