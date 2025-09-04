package main

import "fmt"

type CoffeeMachine interface {
	Brew() string
}

type CapsuleMachine struct {
	Brand string
}

func (c CapsuleMachine) Brew() string {
	return fmt.Sprintf("%s has brewed one cup of coffee", c.Brand)
}

func main() {
	var machine CoffeeMachine
	machine = CapsuleMachine{Brand: "Nespresso"}
	fmt.Println(machine.Brew())
}
