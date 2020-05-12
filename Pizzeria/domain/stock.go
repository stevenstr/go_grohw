/*
* Authod: Stefan
* Created: 05.04.2020
* Last changes: 05.12.2020 22:45
* Task: Class Work Lec4 Project
* This package for stock types
 */

package domain

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const serialSep string = ","

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

//ToString function - convert object has type Stock to simple string from separator
func (s Stock) ToString() string {
	slstring := []string(s)
	return strings.Join(slstring, serialSep)
}

//SaveToFile -
func (s Stock) SaveToFile(filename string) {
	err := ioutil.WriteFile(filename, []byte(s.ToString()), 0666)
	if err != nil {
		fmt.Println("Wrong!", err)
	} else {
		fmt.Println("Success written to", filename, ".")
	}
}

//NewStockFromFile - function has return data from file([]byte) convert it to Stock
func NewStockFromFile(filename string) Stock {
	slbyte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Wrong!", err)
		os.Exit(1)
	} else {
		fmt.Println("Success readed bytes from file ", filename)
	}

	//byte to string
	str := string(slbyte)
	// cut the separator fron string
	//and return []string
	slstring := strings.Split(str, serialSep)

	//convert []string to Stock and return it
	return Stock(slstring)
}
