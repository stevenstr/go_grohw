/*
* Authod: Stefan
* Created: 04.27.2020
* Last changes: 04.27.2020 21:59
* Task: Class Work Lec3
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
