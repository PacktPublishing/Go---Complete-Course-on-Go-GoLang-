package main

import "fmt"

func main() {
	// 15:00 - 17:00
	// isMember
	// orderAmount > 10
	hour := 16
	isMember := true
	orderAmount := 13.50

	// OPTION 1
	if (hour >= 15) && (hour <= 17) && isMember && (orderAmount > 10) {
		fmt.Println("You get 30% off!")
	} else {
		fmt.Println("No Happy Hour deals available")
	}

	// // OPTION 2 (Not recommended)
	// if hour >= 15 {
	// 	if hour <= 17 {
	// 		if isMember {
	// 			if orderAmount > 10 {
	// 				fmt.Println("You get 30% off!")
	// 			} // add also else here
	// 		} // add also else here
	// 	} // add also else here
	// } else {
	// 	fmt.Println("No Happy Hour deals available")
	// }

}
