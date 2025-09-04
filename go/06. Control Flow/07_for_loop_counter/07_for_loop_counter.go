package main

import "fmt"

func main() {
	coffeeCups := 10

	for i := 1; i <= coffeeCups; i++ { // 1, 2, ... 10
		fmt.Printf("Preparing coffee cup #%d\n", i)
	}

	fmt.Println()

	for i := coffeeCups; i >= 1; i-- { // 10, 9, ... 1
		fmt.Printf("Preparing coffee cup #%d\n", i)
	}
}
