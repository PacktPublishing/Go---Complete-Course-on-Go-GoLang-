package main

import "fmt"

func main() {
	var coffeeSizes [3]string // default element values ""
	fmt.Println(coffeeSizes)

	coffeeSizes[0] = "Small"
	fmt.Println(coffeeSizes)

	coffeeSizes[1] = "Medium"
	coffeeSizes[2] = "Large"
	fmt.Println(coffeeSizes)

	coffeeSizes[2] = "Extra Large"
	fmt.Println(coffeeSizes)

	fmt.Println("First element is:", coffeeSizes[0])

	// coffeeSizes[1] = true // compilation error. type is not string
	// coffeeSizes[5] = "Super Extra Large" // index 5 out of bounds
}
