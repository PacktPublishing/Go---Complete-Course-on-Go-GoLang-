package main

import "fmt"

type Barista interface {
	PrepareCoffee() string
}

type SeniorBarista struct {
	Name string
}

func (s SeniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s prepared a caramel latte", s.Name)
}

type JuniorBarista struct {
	Name string
}

func (j JuniorBarista) PrepareCoffee() string {
	return fmt.Sprintf("%s made a hot chocolate", j.Name)
}

func ServeDrink(b Barista) {
	fmt.Println(b.PrepareCoffee())
	fmt.Printf("Barista served coffee to the client\n\n")
}

func main() {
	bogdan := SeniorBarista{Name: "Bogdan"}
	var maria Barista = JuniorBarista{Name: "Maria"}

	ServeDrink(bogdan)
	ServeDrink(maria)

	maria = SeniorBarista{Name: "Maria"}
	ServeDrink(maria)
}
