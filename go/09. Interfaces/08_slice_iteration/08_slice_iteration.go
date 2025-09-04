package main

import "fmt"

type Greeter interface {
	Greet() string
}

type Customer struct {
	Name string
}

func (c Customer) Greet() string {
	return fmt.Sprintf("Customer %s says: Hello! How are you?", c.Name)
}

type Staff struct {
	Role string
}

func (s Staff) Greet() string {
	return fmt.Sprintf("Staff (%s) says: Welcome to the Brew&Beans!", s.Role)
}

func main() {
	greeters := []Greeter{
		Customer{Name: "Bogdan"},
		Staff{Role: "Barista"},
		Customer{Name: "Elena"},
	}

	for _, g := range greeters {
		fmt.Println(g.Greet())
	}
	// "Customer Bogdan says: Hello! How are you?"
	// "Staff (Barista) says: Welcome to the Brew&Beans!"
	// "Customer Elena says: Hello! How are you?"

	fmt.Println()

	greeters = append(greeters, Staff{Role: "Cleaner"})
	for _, g := range greeters {
		fmt.Println(g.Greet())
	}

}
