package main

import "fmt"

type OutOfStockError struct {
	Item string
}

func (e OutOfStockError) Error() string {
	return fmt.Sprintf("%s is out of stock", e.Item)
}

func ServeDrink(item string) (string, error) {
	// Here is your freshly brewed latte/espresso/cappuccino...
	quantity := stock[item]
	if quantity == 0 {
		return "", OutOfStockError{Item: item}
	}
	stock[item]--
	return fmt.Sprintf("Here is your freshly brewed %s", item), nil
}

var stock = map[string]int{
	"Espresso":   5,
	"Latte":      0,
	"Cappuccino": 3,
}

func main() {
	message, err := ServeDrink("Espresso")
	if err != nil {
		fmt.Println("Serving failed!", err)
	} else {
		fmt.Println(message)
	}

	fmt.Println()

	message, err = ServeDrink("Latte")
	if err != nil {
		fmt.Println("Serving failed!", err)
	} else {
		fmt.Println(message)
	}

	fmt.Println()

	message, err = ServeDrink("Tea")
	if err != nil {
		fmt.Println("Serving failed!", err)
	} else {
		fmt.Println(message)
	}

	fmt.Println()

	fmt.Println("Stock after servings:", stock)

}
