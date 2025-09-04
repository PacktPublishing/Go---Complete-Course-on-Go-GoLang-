package main

import "fmt"

type Order struct {
	Customer string
	Item     string
	Quantity int
}

func (o Order) String() string {
	return fmt.Sprintf("Order: %s has ordered %s (%d)", o.Customer, o.Item, o.Quantity)
}

func main() {
	order := Order{Customer: "Bogdan", Item: "Latte", Quantity: 2}
	fmt.Println(order)
}
