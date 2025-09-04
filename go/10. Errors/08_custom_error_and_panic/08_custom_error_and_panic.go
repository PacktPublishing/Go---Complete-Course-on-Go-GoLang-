package main

import "fmt"

type CoffeeError string

func (e CoffeeError) Error() string {
	return string(e)
}

func main() {
	var err error
	err = CoffeeError("No coffee beans loaded!")

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

}
