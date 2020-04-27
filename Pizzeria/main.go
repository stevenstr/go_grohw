/*
* Authod: Stefan
* Created: 04.27.2020
* Last changes: 04.27.2020 18:59
* Task: Class Work Lec3
 */

package main

import (
	"fmt"
)

func main() {
	var pizza string = "Marharita"

	fmt.Println("Value: " + pizza)
	fmt.Println(pizza)
	fmt.Printf("value: %v | type: %T \n\r\n", pizza, pizza)

	var pizza1 string
	pizza1 = "Peperoni"

	fmt.Println("Value: " + pizza1)
	fmt.Println(pizza1)
	fmt.Printf("value: %v | type: %T \n\r\n", pizza1, pizza1)

	pizza2 := "Mashed"
	fmt.Println("Value: " + pizza2)
	fmt.Println(pizza2)
	fmt.Printf("value: %v | type: %T \n\r\n", pizza2, pizza2)
}

func init() {
	fmt.Printf("Project: Pizzeria\nVersion: 0.1\n\r\n")
}
