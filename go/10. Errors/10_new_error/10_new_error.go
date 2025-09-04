package main

import (
	"errors"
	"fmt"
)

func CheckTemperature(temp int) error {
	if temp > 120 {
		return errors.New("Critical failure: temp exceeds safe limit")
	}
	if temp > 90 {
		return fmt.Errorf("Machine overheated at %d grad C", temp)
	}
	return nil
}

func main() {
	temps := []int{75, 95, 130}

	for _, temp := range temps {
		err := CheckTemperature(temp)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("All good. Coffee machine is working fine. Temp is %d\n", temp)
		}
	}

}
