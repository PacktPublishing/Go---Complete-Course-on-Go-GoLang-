package main

import "fmt"

func main() {
	orderTotal := 15.0

	// here condition is true
	if orderTotal > 10 {
		fmt.Println("You get a free cookie with your order!")
	}

	orderTotal = 7.50

	// now condition will be false
	if orderTotal > 10 {
		fmt.Println("You get a free cookie with your order!")
	}

}
