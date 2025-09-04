package main

import (
	"fmt"
	"sync"
	"time"
)

func makeDrink(barista string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Barista %s: Starting to make coffee...\n", barista)
	time.Sleep(2 * time.Second)
	fmt.Printf("Barista %s: Done!\n", barista)
}

func main() {
	var wg sync.WaitGroup
	baristas := []string{"Bogdan", "Elena", "Alex"}

	fmt.Println("Coffee shop opens")

	for _, name := range baristas {
		wg.Add(1)
		go makeDrink(name, &wg)
	}

	wg.Wait()

	fmt.Println("All drinks are ready")
	fmt.Println("Coffee shop closes")
}
