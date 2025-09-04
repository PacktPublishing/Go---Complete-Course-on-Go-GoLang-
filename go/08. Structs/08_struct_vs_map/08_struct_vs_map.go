package main

import "fmt"

type CoffeeOrder struct {
	CustomerName string
	CoffeeType   string
	CoffeeSize   string
}

func main() {
	// Struct version
	orderStruct := CoffeeOrder{
		CustomerName: "Bogdan",
		CoffeeType:   "Latte",
		CoffeeSize:   "Medium",
	}

	// Map version
	orderMap := map[string]string{
		"CustomerName": "Bogdan",
		"CoffeeType":   "Latte",
		"CoffeeSize":   "Medium",
	}

	fmt.Println("---- Using Struct ----")
	fmt.Println("Customer:", orderStruct.CustomerName)
	fmt.Println("Coffee Type:", orderStruct.CoffeeType)
	fmt.Println("Coffee Size:", orderStruct.CoffeeSize)

	fmt.Println()

	fmt.Println("---- Using Map ----")
	fmt.Println("Customer:", orderMap["CustomerName"])
	fmt.Println("Coffee Type:", orderMap["CoffeeType"])
	fmt.Println("Coffee Size:", orderMap["CoffeeSize"])

}
