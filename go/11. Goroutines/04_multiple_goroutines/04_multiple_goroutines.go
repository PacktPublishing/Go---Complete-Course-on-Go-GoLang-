package main

import (
	"fmt"
	"time"
)

func makeDrink(barista string) {
	fmt.Printf("Barista %s: Starting to make coffee...\n", barista)
	time.Sleep(2 * time.Second)
	fmt.Printf("Barista %s: Done!\n", barista)
}

func main() {
	fmt.Println("Coffee shop opens")
	go makeDrink("Bogdan")
	go makeDrink("Elena")
	go makeDrink("Alex")
	time.Sleep(3 * time.Second)

	fmt.Println("All drinks are (probably) ready")
	fmt.Println("Coffee shop closes")
}
