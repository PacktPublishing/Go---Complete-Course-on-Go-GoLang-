package main

import "fmt"

func main() {
	menu := []string{"Espresso", "Latte", "Cake", "Ice cream", "Chocolate"}

	fmt.Println("Today's menu:")

	for index, menuItem := range menu {
		fmt.Printf("%d. %s\n", index+1, menuItem)
	}
}
