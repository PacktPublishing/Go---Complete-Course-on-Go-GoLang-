package main

import "fmt"

func main() {
	dailyIncome := []float64{120.50, 98.75, 140.50, 110.00, 155.50, 170.45, 250.50}
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println("Weekly Income Report:")

	var totalIncome float64 // default zero value is 0.0

	for i, income := range dailyIncome {
		day := daysOfWeek[i]
		fmt.Printf("%s income is $%.2f\n", day, income)
		totalIncome += income
	}

	fmt.Printf("\nTotal weekly income is: $%.0f\n", totalIncome)

	// for _, char := range "Hello from the coffee shop" {
	// 	fmt.Println(char)
	// }
}
