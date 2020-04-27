/*
* Authod: Stefan
* Created: 04.27.2020
* Last changes: 04.27.2020 22:30
* Task: Class Work Lec3 Project
 */

package main

import (
	"fmt"
)

func main() {
	pizza := newPizza()
	fmt.Println(pizza)
}

func newPizza() string {
	return "Margherita"
}

func init() {
	fmt.Printf("Project: Pizzeria\nVersion: 0.1\nLec-3\n\r\n")
}
