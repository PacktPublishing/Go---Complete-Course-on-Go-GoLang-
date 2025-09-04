package main

import "fmt"

func main() {
	stock := make(map[string]int)

	if stock == nil { // in this case with make stock is not nil
		fmt.Printf("Stock map is nil!")
	}

	stock["Espresso"] = 10
	stock["Latte"] = 25

	fmt.Println("Products in stock:", stock)

}
