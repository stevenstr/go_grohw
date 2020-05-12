/*
* Authod: Stefan
* Created: 04.27.2020
* Last changes: 05.12.2020 20:12
* Task: Class Work Lec4 Project
 */

package main

import (
	"fmt"
	"pizzeria/domain"
)

func main() {
	//pizzas now is get value of our constructor
	pizzas := domain.NewStock()

	//and because we can call print method!
	pizzas.Print()
}

//init function - send you the version of project
func init() {
	fmt.Printf("Project: Pizzeria\nVersion: 0.4\nLec-4\n\r\n")
}
