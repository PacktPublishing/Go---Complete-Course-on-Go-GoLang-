package main

import "fmt"

func createTemperatureAdjuster() (func(change float64) float64, float64) {
	baseTemperature := 90.0

	adjustTemperature := func(change float64) float64 {
		baseTemperature = baseTemperature + change
		return baseTemperature
	}

	return adjustTemperature, baseTemperature
}

func main() {
	adjustTemp, originalTemp := createTemperatureAdjuster()
	fmt.Printf("Original temperature is %.1f\n", originalTemp)

	fmt.Printf("Adjusted Temp +1.5: %.1f grad C\n", adjustTemp(1.5))  // baseTemperature is changed
	fmt.Printf("Adjusted Temp -3.0: %.1f grad C\n", adjustTemp(-3.0)) // baseTemperature is changed
	fmt.Printf("Adjusted Temp +5.0: %.1f grad C\n", adjustTemp(5.0))  // baseTemperature is changed

	fmt.Printf("Original temperature is %.1f\n", originalTemp) // originalTemp is not changed
	// 90.0
	// +1.5 -> 91.5
	// -3.0 -> 88.5
	// +5.0 -> 93.5
}
