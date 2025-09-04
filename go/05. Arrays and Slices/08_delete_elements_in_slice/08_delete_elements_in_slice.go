package main

import "fmt"

func deleteByIndex(index int, slice []string) []string {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	coffees := []string{"Espresso", "Latte", "Mocha", "Cappuccino"}

	fmt.Println("Original menu:", coffees)
	fmt.Println("Length is:", len(coffees), "Capacity is:", cap(coffees))

	indexToRemove := 1
	coffees = append(coffees[:indexToRemove], coffees[indexToRemove+1:]...)
	fmt.Println("Updated menu without Latte:", coffees)
	fmt.Println("Length is:", len(coffees), "Capacity is:", cap(coffees))

	indexToRemove = 0

	coffees = deleteByIndex(0, coffees)
	fmt.Println("Updated menu without Espresso:", coffees)
	fmt.Println("Length is:", len(coffees), "Capacity is:", cap(coffees))

}
