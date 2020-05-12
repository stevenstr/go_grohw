/*
* Authod: Stefan
* Created: 05.04.2020
* Last changes: 05.12.2020 20:42
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

//NewStock function - simple constructor
func NewStock() Stock {
	stock := Stock{}

	pizzaNames := []string{
		"Margherita",
		"Cheese",
		"Peperoni",
		"Vegan",
		"GolangPizza",
	}

	diameters := []string{
		"30 cm",
		"40 cm",
	}

	//combine diameter and name of pizza
	for _, pizzaName := range pizzaNames {
		for _, diameter := range diameters {
			stock = append(stock, pizzaName+" "+diameter)
		}
	}

	return stock
}

//GetPackage function -
func GetPackage(s Stock, pkgSize int) (Stock, Stock) {
	return s[:pkgSize], s[pkgSize:]
}
