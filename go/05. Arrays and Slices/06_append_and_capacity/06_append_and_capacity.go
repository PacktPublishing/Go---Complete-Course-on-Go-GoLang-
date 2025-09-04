package main

import "fmt"

func main() {
	menu := []string{"Cake", "Pie"}

	fmt.Println("Initial menu:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) // 2, 2
	fmt.Printf("Memory location: %p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])

	fmt.Println()

	menu = append(menu, "Donut")
	fmt.Println("Menu after adding donut:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) // 3, 4
	fmt.Printf("Memory location: %p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])

	fmt.Println()

	menu = append(menu, "Ice cream")
	fmt.Println("Menu after adding ice cream:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) // 4, 4
	fmt.Printf("Memory location: %p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])

	fmt.Println()

	menu = append(menu, "Cream")
	fmt.Println("Menu after adding cream:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu)) //
	fmt.Printf("Memory location: %p\n", &menu)
	fmt.Printf("Memory location of \"Cake\": %p\n", &menu[0])

	fmt.Println()

	cupSizes := make([]string, 0, 5) // len will be 0 and capacity will be 5
	fmt.Println("Len of cupSizes:", len(cupSizes), "Capacity of cupSizes:", cap(cupSizes))
	// cupSizes[0] = "Small" // panic: runtime error: index out of range [0] with length 0
	cupSizes = append(cupSizes, "Small", "Medium") // len becomes 2
	fmt.Println("Len of cupSizes:", len(cupSizes), "Capacity of cupSizes:", cap(cupSizes))
	cupSizes[0] = "Extra small"
	fmt.Println(cupSizes) // no error now, len is 2
	fmt.Println("Len of cupSizes:", len(cupSizes), "Capacity of cupSizes:", cap(cupSizes))

}
