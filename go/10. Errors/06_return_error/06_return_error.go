package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("coffee_orders.txt")

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File doesn't exist")
		} else {
			fmt.Println("General file opening error", err)
		}
		return
	}

	fmt.Println("Successfully accessed file:", file.Name())

}
