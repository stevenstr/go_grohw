/*
* Author: Stefan
* Last changes: 05.27.2020 18:45
* Task: Workshop L4
 */

package domain

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

const sep = ","

//Stock slice
type Stock []string

//NewStock func
func NewStock() Stock {
	stock := Stock{}
	listpzz := []string{"Margaritta", "Cheese", "Onionworld", "Strawberry"}
	diameter := []string{"30 cm", "40 cm"}

	for _, pizza := range listpzz {
		for _, diameter := range diameter {
			stock = append(stock, pizza+" "+diameter)
		}
	}
	return stock
}

//Print method
func (s Stock) Print() {
	fmt.Println("Print stock!")
	for i, v := range s {
		fmt.Println(i, v)
	}
}

//Sort method
func (s Stock) Sort() Stock {
	fmt.Println("Sort:")
	slstr := []string(s)
	sort.Strings(slstr)
	return Stock(slstr)
}

//SaveToFile method
func (s Stock) SaveToFile(filename string) {
	err := ioutil.WriteFile(filename, []byte(s.ToString()), 0666)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success writen!")
	}
}

//ToString method
func (s Stock) ToString() string {
	sl := []string(s)
	str := strings.Join(sl, sep)
	return str
}

//NewStockFromFile function
func NewStockFromFile(filename string) Stock {
	slbyte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Read from File!")
	str := string(slbyte)
	slstr := strings.Split(str, sep)
	return Stock(slstr)
}
