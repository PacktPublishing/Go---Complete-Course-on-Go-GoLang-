package main

import (
	"fmt"
	"strings"
)

type MenuList []string

func (ml MenuList) String() string {
	// [Coffee, Tea, Croissant]
	// OPTION 1
	return fmt.Sprintf("[%s]", strings.Join(ml, ", "))

	// // OPTION 2
	// return "[" + strings.Join(ml, ", ") + "]"

	// // OPTION 3
	// c := "["
	// for i, mi := range ml {
	// 	c += mi
	// 	if i < len(ml)-1 {
	// 		c += ", "
	// 	}
	// }
	// c += "]"
	// return c
}

func main() {
	menu := MenuList{"Coffee", "Tea", "Croissant"}
	fmt.Println("Menu:", menu)
}
