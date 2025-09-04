package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType   string // ""
	CoffeeSize   string // ""
	CustomerName string // ""
	BonusPoints  int    // 0
}

func main() {
	var order CoffeeOrder

	fmt.Println(order) // {   0}

	order.CoffeeType = "Cappuccino"
	order.CoffeeSize = "Large"
	order.CustomerName = "Bogdan"
	order.BonusPoints = 15

	// order.IsReady = true // order.IsReady undefined (type CoffeeOrder has no field or method IsReady)

	fmt.Println(order) // {Cappuccino Large Bogdan 15}
	fmt.Println("Coffee type for the order is:", order.CoffeeType)
}
