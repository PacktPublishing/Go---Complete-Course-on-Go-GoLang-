package main

import "fmt"

// func estimateBrewTime(cupsQty int, secondsPerCup int) int {
// 	totalTimeSeconds := cupsQty * secondsPerCup
// 	return totalTimeSeconds
// }

func estimateBrewTime(cupsQty int, secondsPerCup int) (totalTimeSeconds int, info string) {
	totalTimeSeconds = cupsQty * secondsPerCup
	info = "Estimated total brew time:"
	return // naked return
}

func main() {
	// 12 cups, 20 seconds per cup
	// 12 * 20 = 240 seconds
	brewTime, info := estimateBrewTime(12, 20)
	fmt.Println(info, brewTime)
}
