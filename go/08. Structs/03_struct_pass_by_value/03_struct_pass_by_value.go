/*
* List of types in Go which are passed by value:
- Basic types: int, float32, float64, bool, string, complex64
- Arrays
- Structs
- Pointers
*
*/
package main

import "fmt"

type CoffeeOrder struct {
	CoffeeType string
	Size       string
}

func changeCoffeeSize(order CoffeeOrder) { // copy value of the myOrder
	order.Size = "Small" // Attempt to modify order in the function
	fmt.Println("Order size in the function:", order.Size)
}

func main() {
	myOrder := CoffeeOrder{"Cappuccino", "Large"}
	fmt.Println("My order in the main function before changeCoffeeSize function call:", myOrder)

	changeCoffeeSize(myOrder) // change of the size in the changeCoffeeSize doesn't affect myOrder
	fmt.Println("My order in the main function after changeCoffeeSize function call:", myOrder)
}
