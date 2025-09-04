package main

import (
	"fmt"
	"time"
)

func makeDrink(barista string, drink string, c chan string) {
	fmt.Printf("Barista %s: Starting to make a drink...\n", barista)
	time.Sleep(2 * time.Second)
	msg := fmt.Sprintf("Barista %s: %s is ready!", barista, drink)
	c <- msg
}

func main() {
	fmt.Println("Coffee shop opens...")

	c := make(chan string)

	baristas := []string{"Bogdan", "Elena", "Alex"}
	drinks := []string{"Latte", "Espresso", "Tea"}

	for i, barista := range baristas {
		go makeDrink(barista, drinks[i], c)
	}

	for range baristas {
		msg := <-c
		fmt.Println(msg)
	}

	fmt.Println("All drinks are ready!")
	fmt.Println("Coffee shop closes...")

}
