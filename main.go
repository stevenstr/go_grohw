/*
* Author: Stefan
* Last changes: 05.27.2020 18:45
* Task: Workshop L4
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"pizza/domain"
	"strings"
)

func main() {

	pizzas := domain.NewStock()
	pizzas.Print()
	pizzas.Sort()
	pizzas.Print()

	filename := getFileName()
	pizzas.SaveToFile(filename)

	pizzasRead := domain.NewStockFromFile(filename)
	pizzasRead.Print()
	pizzasRead.Sort()
	pizzasRead.Print()
}

func getFileName() string {
	in := bufio.NewReader(os.Stdin)
	fn, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Error:", err)
	}
	fn = strings.TrimSpace(fn)
	return fn
}
