package main

import "fmt"

func main() {
	var err error
	err = fmt.Errorf("Some interesting coffee machine error")

	// err = "Some interesting coffee machine error" // string does not implement error (missing method Error)

	if err == nil {
		fmt.Println("There is no error")
	} else {
		fmt.Println("Error occurred!", err)
	}

}
