/*
* Author: Stefan
* Craete: 04.22.2020 11:25
* Last changes:04.22.2020 12:23
* Task: Class Work Lec1-2
 */

package main

import (
	"fmt"
	"module/helper"
)

func main() {
	fmt.Println("Hello golang world!!!")
	helper.Core()
	helper.Help()

	//exported
	//helper.Help()

	//unexported
	//helper.elp()
}
