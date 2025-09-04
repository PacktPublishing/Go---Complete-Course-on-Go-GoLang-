package main

import "fmt"

func main() {
	day := "Sunday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday special: Buy one get one 50% off")
	case "Saturday", "Sunday":
		fmt.Println("Weekend special: Free croissant with any coffee!")
	default:
		fmt.Println("Unknown day")
	}

	// if day == "Monday" || day == "Tuesday" {
	// 	fmt.Println("Weekday special: Buy one get one 50% off")
	// } else if day == "Saturday" || day == "Sunday" {
	// 	fmt.Println("Weekend special: Free croissant with any coffee!")
	// } else {
	// 	fmt.Println("Unknown day")
	// }

}
