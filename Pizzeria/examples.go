package main

import "fmt"

func examples() {
	//ex 1 Variables definition

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

	//ex 7 Constants 1
	/*
		//Pi const
		const Pi = 3.14
		//World const
		const World = "世界"

		fmt.Println("Hello,", World)
		fmt.Println("Happy", Pi, "Day!!!")
		//Truth const
		const Truth = true
		fmt.Println("Go rules?", Truth)
	*/

	//ex 8 Constatns 2
	/*
		//full form
		const (
			//Nature const
			Nature = iota
			//Science const
			Science = iota
			//Social const
			Social = iota
		)

		fmt.Println(Nature)
		fmt.Println(Science)
		fmt.Println(Social)

		//short form
		const (
			//Nature1 const
			Nature1 = iota
			//Science1 const
			Science1
			//Social1 const
			Social1
		)

		fmt.Println(Nature1)
		fmt.Println(Science1)
		fmt.Println(Social1)
	*/

	//ex 9 Type convertion
	/*
		var x, y int = 3, 4
		var f float64 = math.Sqrt(float64(x*x + y*y))
		var z uint = uint(f)

		fmt.Printf("value x: %v | type x: %T \n\r\n", x, x)
		fmt.Printf("value y: %v | type y: %T \n\r\n", y, y)
		fmt.Printf("SQRT value: %v | type: %T \n\r\n", f, f)
		fmt.Printf("SQRT value: %v | type: %T \n\r\n", z, z)

		fmt.Print(x, y, z)
	*/

	//ex 10 Type conversation 1
	/*
		var i int = 1
		var j float64 = float64(i)
		var f uint = uint(j)

		fmt.Printf("value: %v | type: %T \n\r\n", i, i)
		fmt.Printf("value: %v | type: %T \n\r\n", j, j)
		fmt.Printf("value: %v | type: %T \n\r\n", f, f)
	*/

	//ex 11 Type conversation 2
	/*
		var foo = 412
		var bar = "-100"

		var resstrig string
		var resint int

		resstrig = strconv.Itoa(foo)
		fmt.Println("From number", foo, "to string", resstrig)

		resint, err := strconv.Atoi(bar)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("From string", bar, "to int", resint)
	*/
}
