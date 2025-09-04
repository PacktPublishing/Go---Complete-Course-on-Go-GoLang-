package main

import "fmt"

type CoffeeMachine interface {
	Brew() string
}

// CapsuleMachine implementation
type CapsuleMachine struct {
	Brand string
	Model string
	Price int
}

func (c CapsuleMachine) Brew() string {
	return fmt.Sprintf("%s %s has brewed a cup of capsule coffee", c.Brand, c.Model)
}

// DripMachine implementation
type DripMachine struct {
	Model string
	Price int
}

func (d DripMachine) Brew() string {
	return fmt.Sprintf("Drip coffee shot was prepared by %s", d.Model)
}

func (d DripMachine) DeepClean() {
	fmt.Println("Deep cleaning of the", d.Model)
}

func main() {
	var machineOne CoffeeMachine
	var machineTwo CoffeeMachine

	machineOne = CapsuleMachine{"Nespresso", "XB23", 135}
	machineTwo = DripMachine{"BrewPro", 250}

	fmt.Println(machineOne.Brew())
	fmt.Println(machineTwo.Brew())
	// machineTwo.DeepClean() // !!! Not possible because type of machineTwo is CoffeeMachine

	var machineThree DripMachine
	machineThree = DripMachine{"SuperPowerDrip", 470}
	machineThree.DeepClean() // !!! Here is possible to call because now it has type DripMachine

	machineTwo = CapsuleMachine{"New Brand", "324", 760} // can reassign value of other type because of interface
	fmt.Println(machineTwo.Brew())

}
