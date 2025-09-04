package main

import "fmt"

type AnyValue interface{} // is not necessary to create custom type

func LogAnyValue(v interface{}) {
	fmt.Println(v)
}

// any is an alias for interface{} and is equivalent to interface{} in all ways.
func LogAnyValueWithAny(v any) {
	fmt.Println(v)
}

func main() {
	// can assign value of any type
	var any AnyValue = "Coffee"
	fmt.Println(any)

	any = 10
	fmt.Println(any)

	any = []string{"Latte", "Espresso"}
	fmt.Println(any)

	// can assign value of any type
	var anotherAny interface{} = "Latte"
	anotherAny = 10.5
	anotherAny = true
	fmt.Println(anotherAny)

	// slice accepts values of any types
	valuesOfDifferentTypes := []interface{}{
		"Latte",
		50.5,
		true,
		[3]int{10, 5, 3},
	}

	for _, e := range valuesOfDifferentTypes {
		fmt.Println(e)
	}

	// Call a function with any value
	LogAnyValue("Bogdan")
	LogAnyValue(true)
	LogAnyValueWithAny([2]string{"Latte", "Espresso"})
}
