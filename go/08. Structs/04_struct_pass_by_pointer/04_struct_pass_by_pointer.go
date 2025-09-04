package main

import "fmt"

type CoffeeMachine struct {
	Model          string
	Status         string
	OperationHours int
}

func markAsOutOfService(machine *CoffeeMachine) {
	machine.Status = "Out of Service"
	fmt.Println("In the function - Machine status changed to:", machine.Status)
}

func main() {
	espressoMachine := CoffeeMachine{
		Model:          "Extra class Espresso machine 234A",
		Status:         "Operational",
		OperationHours: 75,
	}

	pointerToEspressoMachine := &espressoMachine

	markAsOutOfService(pointerToEspressoMachine)
	fmt.Println("Machine status in the main function:", espressoMachine.Status)
}
