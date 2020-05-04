/*
* Authod: Stefan
* Created: 04.27.2020
* Last changes: 05.04.2020 20:50
* Task: Class Work Lec4 Project
 */

package main

import (
	"fmt"
	"pizzeria/domain"
)

func main() {
	//pizzas now is variable of Stock type
	pizzas := domain.Stock{newPizza(), "Cheese"}
	pizzas = append(pizzas, "Peperoni")
	pizzas = append(pizzas, "Vegan")
	pizzas = append(pizzas, "GolangPizza")
	//and because we can call print method!
	pizzas.Print()
}

//add first pizzas
func newPizza() string {
	return "Margherita"
}

//init function - send you the version of project
func init() {
	fmt.Printf("Project: Pizzeria\nVersion: 0.3\nLec-4\n\r\n")
}
