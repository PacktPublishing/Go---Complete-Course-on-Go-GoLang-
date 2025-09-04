package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType string
	Size       string
}

func main() {
	// struct literal
	order := CoffeeOrder{
		CoffeeType: "Latte",
		Size:       "Large",
	}

	// inline struct literal
	myFriendsOrder := CoffeeOrder{"Espresso", "Small"}

	fmt.Println(order)
	fmt.Println("Type of coffee in the order:", order.CoffeeType)
	fmt.Println("Size of coffee in the order:", order.Size)

	fmt.Println()

	fmt.Println(myFriendsOrder)
	fmt.Println("Type of coffee in the myFriendsOrder:", myFriendsOrder.CoffeeType)
	fmt.Println("Size of coffee in the myFriendsOrder:", myFriendsOrder.Size)

}
