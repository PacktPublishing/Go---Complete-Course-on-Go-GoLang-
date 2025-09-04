package main

import "fmt"

func main() {
	dessertMenu := [4]string{"Muffin", "Brownie", "Croissant", "Cookie"}
	fmt.Println("Dessert Menu:", dessertMenu)

	slice := dessertMenu[1:3] // elements with indexes 1, 2
	fmt.Println("Slice of the Dessert Menu [1:3]:", slice)

	slice = dessertMenu[:] // all elements
	fmt.Println("Slice of the Dessert Menu [:]:", slice)

	slice = dessertMenu[2:] // all elements starting from index 2
	fmt.Println("Slice of the Dessert Menu [2:]:", slice)

	slice = dessertMenu[:3] // all elements from the beginning to index 3 (not inclusive)
	fmt.Println("Slice of the Dessert Menu [:3]:", slice)

}
