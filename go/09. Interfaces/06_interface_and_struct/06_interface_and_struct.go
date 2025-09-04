package main

import "fmt"

type Describable interface {
	Description() string
}

type Tea struct {
	Type string
	Size string
}

func (t Tea) Description() string {
	return fmt.Sprintf("A %s cup of %s tea.", t.Size, t.Type)
}

func main() {
	var d Describable = Tea{Type: "Green", Size: "Large"}
	fmt.Println(d.Description()) // "A Large cup of Green tea."
}
