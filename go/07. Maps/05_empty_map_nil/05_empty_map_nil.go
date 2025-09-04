package main

import "fmt"

func main() {
	// var drink string // zero value is ""
	var stock map[string]int // zero value is nil
	fmt.Println(stock)
	fmt.Printf("Location in memory %p\n", &stock)

	if stock == nil {
		fmt.Printf("Stock map is nil!")
	}

	stock["Espresso"] = 10 // Stock map is nil!panic: assignment to entry in nil map

	fmt.Println("Products in stock:", stock)

}
