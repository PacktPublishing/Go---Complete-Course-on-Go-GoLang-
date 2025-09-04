package main

import "fmt"

func main() {
	points := 75

	// more specific conditions have to be higher
	if points >= 100 { // [100, ...]
		fmt.Println("Platinum member: Free coffee every day!")
	} else if points >= 50 { // [50, 100)
		fmt.Println("Gold member: 20% discount on latte")
	} else if points >= 20 { // [20, 50)
		fmt.Println("Silver member: Free cookie on Monday")
	} else { // [0, 20)
		fmt.Println("Bronze member: Keep sipping to earn rewards")
	}

}
