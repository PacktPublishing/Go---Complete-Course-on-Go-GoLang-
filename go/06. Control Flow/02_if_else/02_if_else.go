package main

import "fmt"

func main() {
	customerName := "Bogdan"
	orderedSize := "Large"

	if orderedSize == "Small" {
		fmt.Println(customerName, "ordered a Small coffee. It will be ready in 2 min")
	} else {
		fmt.Println(customerName, "ordered something bigger. It might take a bit more time")
	}

}
