/*
* Authod: Stefan
* Created: 05.04.2020
* Last changes: 05.04.2020 20:50
* Task: Class Work Lec4 Project
* This package for stock types
 */

package domain

import "fmt"

//Stock type
type Stock []string

//Print method - get copy of variable with Stock type, and can use her value called name stock
func (s Stock) Print() {
	fmt.Println("List of pizzas...")
	for index, pizza := range s {
		fmt.Println("id: ", index, " | pizza: ", pizza)
	}
}
