/*
* Authod: Stefan
* Created: 05.04.2020
* Last changes: 05.12.2020 21:40
* Task: Class Work Lec4 OOP in Go
 */

package main

import (
	"fmt"
)

func examples() { //examples

	//ex 1 Switch
	/*
		fmt.Print("Go runs on ")
		switch os := runtime.GOOS; os {
		case "darvin":
			fmt.Printf("OS X\n")
		case "linux":
			fmt.Printf("GNU/Linux\n")
		default:
			fmt.Printf("%s\n", os)
		}
	*/

	//ex 2 Switch 2
	/*
		fmt.Println("When's saturday?")
		today := time.Now().Weekday()
		switch time.Saturday {
		case today + 0:
			fmt.Println("today")
		case today + 1:
			fmt.Println("tomorrow")
		case today + 2:
			fmt.Println("in two days")
		default:
			fmt.Println("too far away")
		}
	*/

	//ex 3 Switch 3 in fact as switch true
	/*
		t := time.Now()

		switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon!")
		default:
			fmt.Println("Good evening!")
		}
	*/

	//ex 5 Labels/Marks goto
	/*
			fmt.Println(1)
			goto End
			fmt.Println(2) // unreachable code
		End:
			fmt.Println(3)
	*/

	//ex 6 Switch 4 fallthrough
	/*
		fmt.Println("When's saturday?")
		today := time.Now().Weekday()
		switch time.Saturday {
		case today + 0:
			fmt.Println("today")
		case today + 1:
			fmt.Println("tomorrow")
		case today + 2:
			fmt.Println("in two days")
		case today + 3:
			fmt.Println("in three days")
			fallthrough
		case today + 4:
			fmt.Println("in four days")
		case today + 5:
			fmt.Println("in five days")
			fallthrough
		default:
			fmt.Println("too far away")
		}
	*/

	//ex 7 break
	/*
		for i := 0; i < 10; i++ {
			f()
		}
	*/

	//ex 8 break and marks
	/*
		OuterLoop:
			for i := 0; i < 10; i++ {
				fmt.Println(i)
				for j := 0; j < 10; j++ {
					fmt.Println(j)
					break OuterLoop
				}
			}
	*/

	//ex 9 continue and marks
	/*
	   OuterLoop:
	   	for i := 0; i < 10; i++ {
	   		fmt.Println("i= ", i)
	   		for j := 0; j < 10; j++ {
	   			fmt.Println("j= ", j)
	   			continue OuterLoop
	   		}
	   	}
	*/

	//ex 10 defer
	/*
		defer fmt.Println(", world!")
		defer fmt.Println(", world!1")
		defer fmt.Println(", world!2")
		fmt.Print("Hello")
	*/

	//ex 11 defer
	/*
		s := "Hello"
		defer fmt.Println(s)

		s = "bye"
		fmt.Println(s)
	*/

	//ex 12 defer stack | runtime: 0 1 2 3 | defer: 3 2 1 0
	/*
		for i := 0; i < 4; i++ {
			fmt.Println("runtime out:", i)
			defer fmt.Println("defer out:", i)
		}
		fmt.Println("done")
	*/

	//ex 13 slice byte

	sl := []byte("Hello")
	fmt.Println(sl)
}

//ex 7
/*
func f() {
	break
}
*/
