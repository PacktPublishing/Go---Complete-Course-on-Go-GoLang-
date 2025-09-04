package main

import "fmt"

func divideByZero(a int) {
	fmt.Println(10 / a)
}

func barista() {
	fmt.Println("Barista: Starting to make coffee...")
	fmt.Println("Barista: Done!")
	divideByZero(0)
}

func main() {
	fmt.Println("Coffee shop opens")
	barista()
	fmt.Println("Coffee shop closes")
}
