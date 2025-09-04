package main

import "fmt"

func main() {
	desserts := [3]string{"Cupcake", "Eclair", "Ice cream"}
	fmt.Println("Array:", desserts)

	fmt.Println()

	slice := desserts[:1] // cap will be the same as original array len
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)

	fmt.Println()

	slice = append(slice, "Macaron")
	fmt.Println("Array:", desserts)
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)

	fmt.Println()

	slice = append(slice, "Cake") // len will become 3 and cap is 3
	fmt.Println("Array:", desserts)
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)

	fmt.Println()

	// because len is already equal to cap -> new array is allocated for the slice
	slice = append(slice, "Juice")
	fmt.Println("Array:", desserts)
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)

	fmt.Println()
	slice[0] = "Chocolate" // original array is not modified
	fmt.Println("Array:", desserts)
	fmt.Println("Len of slice:", len(slice), "Cap of slice:", cap(slice), "Slice:", slice)
}
