/*
* Authod: Stefan
* Created: 04.27.2020
* Last changes: 05.02.2020 17:26
* Task: Class Work Lec3 Project
 */

package main

import (
	"fmt"
)

func main() {
	pizzas := []string{newPizza(), "Cheese"}
	pizzas = append(pizzas, "Peperoni")
	pizzas = append(pizzas, "Vegan")
	pizzas = append(pizzas, "GolangPizza")

	for index, pizza := range pizzas {
		fmt.Println(index, pizza)
	}
}

//add first pizzas
func newPizza() string {
	return "Margherita"
}

//init function - send you the version of project
func init() {
	fmt.Printf("Project: Pizzeria\nVersion: 0.2\nLec-3\n\r\n")
}
