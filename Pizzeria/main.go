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
	//ex 1 Variables definition
	/*
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
	*/
	// ex 2
	/*
		var c, java, python bool

		fmt.Printf("value: %v | type: %T \n\r\n", c, c)
		fmt.Printf("value: %v | type: %T \n\r\n", java, java)
		fmt.Printf("value: %v | type: %T \n\r\n", python, python)

		bb, cc, dd := true, 1.2, "DDDDD"

		fmt.Printf("value: %v | type: %T \n\r\n", bb, bb)
		fmt.Printf("value: %v | type: %T \n\r\n", cc, cc)
		fmt.Printf("value: %v | type: %T \n\r\n", dd, dd)
	*/
	//ex 3 Zero Values
	/*
		//ints
		var i int
		var i8 int8
		var i16 int16
		var i32 int32
		var i64 int64

		fmt.Printf("Ints \n\r\n")
		fmt.Printf("value: %v | type: %T \n\r\n", i, i)
		fmt.Printf("value: %v | type: %T \n\r\n", i8, i8)
		fmt.Printf("value: %v | type: %T \n\r\n", i16, i16)
		fmt.Printf("value: %v | type: %T \n\r\n", i32, i32)
		fmt.Printf("value: %v | type: %T \n\r\n", i64, i64)

		//uints
		var ui uint
		var ui8 uint8
		var ui16 uint16
		var ui32 uint32
		var ui64 uint64

		fmt.Printf("Uints \n\r\n")
		fmt.Printf("value: %v | type: %T \n\r\n", ui, ui)
		fmt.Printf("value: %v | type: %T \n\r\n", ui8, ui8)
		fmt.Printf("value: %v | type: %T \n\r\n", ui16, ui16)
		fmt.Printf("value: %v | type: %T \n\r\n", ui32, ui32)
		fmt.Printf("value: %v | type: %T \n\r\n", ui64, ui64)

		//floats

		var fl32 float32
		var fl64 float64

		fmt.Printf("Floats \n\r\n")
		fmt.Printf("value: %v | type: %T \n\r\n", fl32, fl32)
		fmt.Printf("value: %v | type: %T \n\r\n", fl64, fl64)

		//bools

		var bootr bool = true
		var boofl bool

		fmt.Printf("Bools \n\r\n")
		fmt.Printf("value: %v | type: %T \n\r\n", bootr, bootr)
		fmt.Printf("value: %v | type: %T \n\r\n", boofl, boofl)
	*/
	//ex 4
	/*
		var a, b = 12, "sheep"
		var foo, bar int16 = 1, 0
		c, d := "Shhhh", "dog"

		fmt.Printf("value: %v | type: %T \n\r\n", a, a)
		fmt.Printf("value: %v | type: %T \n\r\n", b, b)
		fmt.Printf("value: %v | type: %T \n\r\n", foo, foo)
		fmt.Printf("value: %v | type: %T \n\r\n", bar, bar)
		fmt.Printf("value: %v | type: %T \n\r\n", c, c)
		fmt.Printf("value: %v | type: %T \n\r\n", d, d)
	*/

	//ex 5 fmt
	/*
		var a int
		var st string
		var boo bool
		var flo float64
		var compl64 complex64

		fmt.Printf("value: %v | type: %T \n\r\n", a, a)
		fmt.Printf("value: %v | type: %T \n\r\n", st, st)
		fmt.Printf("value: %v | type: %T \n\r\n", boo, boo)
		fmt.Printf("value: %v | type: %T \n\r\n", flo, flo)
		fmt.Printf("value: %v | type: %T \n\r\n", compl64, compl64)
	*/

	//ex 6 Constants
	/*
		const (
			a = 1
			b	//also 1
			c = "test"
			d	//also "test"
		)

		fmt.Printf("value: %v | type: %T \n\r\n", a, a)
		fmt.Printf("value: %v | type: %T \n\r\n", b, b)
		fmt.Printf("value: %v | type: %T \n\r\n", c, c)
		fmt.Printf("value: %v | type: %T \n\r\n", d, d)

		const ges int = 1888

		fmt.Printf("value: %v | type: %T \n\r\n", ges, ges)

		const (
			jan = iota // iota is a counter start with zero
			feb
			mar
			aip
			may
		)

		fmt.Printf("value: %v | type: %T \n\r\n", jan, jan) //0
		fmt.Printf("value: %v | type: %T \n\r\n", feb, feb) //1
		fmt.Printf("value: %v | type: %T \n\r\n", mar, mar) //2
		fmt.Printf("value: %v | type: %T \n\r\n", aip, aip) //3
		fmt.Printf("value: %v | type: %T \n\r\n", may, may) //4
	*/
	//ex 7

	const Pi = 3.14
	const World = "世界"

	fmt.Println("Hello,", World)
	fmt.Println("Happy", Pi, "Day!!!")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func init() {
	fmt.Printf("Project: Pizzeria\nVersion: 0.1\n\r\n")
}
