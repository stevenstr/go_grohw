/*
* Authod: Stefan
* Created: 04.27.2020
* Last changes: 05.02.2020 18:00
* Task: Class Work Lec3
 */

package main

func examples() { //examples
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

	//ex 12 Functions
	/*
		fmt.Println(add(3, 4))
		fmt.Println(add(5, 4))
		fmt.Println(add(1, 4))
		fmt.Println(add(9, 4))
	*/

	//ex 13 Functions 1
	/*
		fmt.Println(add(3, 4))
		fmt.Println(add(5, 4))
		fmt.Println(add(1, 4))
		fmt.Println(add(9, 4))
	*/

	//ex 14 Functions 2
	/*
		x, y := "world!", "Hello, "
		a, b := swap(x, y)
		fmt.Println("Orig:", x, y)
		fmt.Println("Swaps:", a, b)
	*/

	//ex 15 Functions 3
	/*
		a, _ := swap("a", "b") //b
		fmt.Println(a)

		_, b := swap("a", "b") //a
		fmt.Println(b)
	*/

	//ex 16 Functions 4
	/*
		fmt.Println(split(17))
	*/

	//ex 17 Functions 5 Function as a type
	/*
		hypot := func(x, y float64) float64 {
			return math.Sqrt(x*x + y*y)
		}
		fmt.Println(hypot(5, 22))

		fmt.Println(compute(hypot))
		fmt.Println(compute(math.Pow))
	*/

	//ex 18 Closure
	/*
		pos, neg := adder(), adder()
		for i := 0; i < 10; i++ {
			fmt.Println(
				pos(i),
				neg(-2*i),
			)
		}
	*/

	//ex 19 Array
	/*
		var a [5]int
		var b = [5]int{1, 2, 3, 4, 5}
		fmt.Println(a)
		fmt.Println(b)

		c := [4]int{1, 2, 3, 4}
		fmt.Println(c)

		foo := [2]string{"Hello,", "world!"}
		fmt.Println(foo, foo[0], foo[1])

		var bar [2]string
		fmt.Println(bar)

		bar[0] = "Golang it's"
		bar[1] = "my life..."
		fmt.Println(bar)
		fmt.Println(bar[0])
		fmt.Println(bar[1])
	*/

	//ex 20 Slice
	/*
		var sl []int
		var sl1 = []int{1, 2}
		var sl2 = []int{}
		sl3 := []int{1, 2, 3}

		fmt.Println(sl)
		fmt.Println("sl is nil?", sl == nil)
		fmt.Println(sl1)
		fmt.Println("sl1 is nil?", sl1 == nil)
		fmt.Println(sl2)
		fmt.Println("sl2 is nil?", sl2 == nil)
		fmt.Println(sl3)
		fmt.Println("sl3 is nil?", sl3 == nil)
	*/

	//ex 21 Slice 1
	/*
		arr := [5]int{1, 2, 3, 4, 5}
		fmt.Println("arr:", arr)
		fmt.Printf("type: %T | len: %v\n\r\n", arr, len(arr))

		sl := arr[:] //1, 2, 3, 4, 5
		fmt.Println("sl:", sl)
		fmt.Printf("type: %T | len: %v | cap: %v\n\r\n", sl, len(sl), cap(sl))

		sl1 := arr[1:4] //2, 3, 4
		fmt.Println("sl1", sl1)
		fmt.Printf("type: %T | len: %v | cap: %v\n\r\n", sl1, len(sl1), cap(sl1))

		sl1 = append(sl1, 1)
		fmt.Println("sl1", sl1)
		fmt.Printf("type: %T | len: %v | cap: %v\n\r\n", sl1, len(sl1), cap(sl1))
		fmt.Println("sl:", sl)
		fmt.Println("arr:", arr)
	*/

	//ex 22 Slice 2
	/*
		arr := [7]int{0, 1, 2, 3, 4, 5, 6}
		fmt.Println("arr:", arr, len(arr))

		sl := arr[1:5]
		fmt.Println("arr:", arr, len(arr))
		fmt.Println("sl:", sl, len(sl), cap(sl))
		fmt.Println()

		sl = sl[:cap(sl)]
		fmt.Println("arr:", arr, len(arr))
		fmt.Println("sl:", sl, len(sl), cap(sl))
		fmt.Println()
	*/

	//ex 23 Slice 3
	/*
		nam := []string{
			"Ivan",
			"Kolya",
			"Lesha",
			"Sergio",
		}

		sl1 := nam[1:4]
		sl2 := nam[0:3]
		sl3 := nam[2:4]

		fmt.Printf("arr: %v | type: %T | len: %v \n\r\n", nam, nam, len(nam))
		fmt.Printf("sl1: %v | type: %T | len: %v | cap: %v\n\r\n", sl1, sl1, len(sl1), cap(sl1))
		fmt.Printf("sl2: %v | type: %T | len: %v | cap: %v\n\r\n", sl2, sl2, len(sl2), cap(sl2))
		fmt.Printf("sl3: %v | type: %T | len: %v | cap: %v\n\r\n", sl3, sl3, len(sl3), cap(sl3))

		fmt.Println("==================================================")
		sl2[1] = "Junk" //side effect because cap has not reached its limit

		fmt.Printf("arr: %v | type: %T | len: %v \n\r\n", nam, nam, len(nam))
		fmt.Printf("sl1: %v | type: %T | len: %v | cap: %v\n\r\n", sl1, sl1, len(sl1), cap(sl1))
		fmt.Printf("sl2: %v | type: %T | len: %v | cap: %v\n\r\n", sl2, sl2, len(sl2), cap(sl2))
		fmt.Printf("sl3: %v | type: %T | len: %v | cap: %v\n\r\n", sl3, sl3, len(sl3), cap(sl3))

		sl2 = append(sl2, "Junk1", "Junk2", "Junk3")
		fmt.Println("==================================================")
		//cap was reached and overreached, and sl2 is NEW

		fmt.Printf("arr: %v | type: %T | len: %v \n\r\n", nam, nam, len(nam))
		fmt.Printf("sl1: %v | type: %T | len: %v | cap: %v\n\r\n", sl1, sl1, len(sl1), cap(sl1))
		fmt.Printf("sl2: %v | type: %T | len: %v | cap: %v\n\r\n", sl2, sl2, len(sl2), cap(sl2))
		fmt.Printf("sl3: %v | type: %T | len: %v | cap: %v\n\r\n", sl3, sl3, len(sl3), cap(sl3))

		fmt.Println("==================================================")
	*/

	//ex 24 Slice 4
	/*
		q := []int{2, 3, 5, 7, 11, 13}
		fmt.Println(q)

		r := []bool{true, false, true, true, false, true}
		fmt.Println(r)

		s := []struct {
			i int
			b bool
		}{
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
			{13, true},
		}
		fmt.Println(s)
	*/

	//ex 25 Slice 5
	/*
		s := []int{3, 5, 7, 9, 11}
		fmt.Println(s, len(s), cap(s))
		s = s[1:4]
		fmt.Println(s, len(s), cap(s))
		s = s[:2]
		fmt.Println(s, len(s), cap(s))
		s = s[1:]
		fmt.Println(s, len(s), cap(s))
	*/

	//ex 26 Slice 6 Len Cap Grows
	/*
		sl := []int{}

		for i := 0; i < 45; i++ {
			fmt.Printf("len: %v | cap: %v\n\r", len(sl), cap(sl))
			sl = append(sl, i)
		}
	*/

	//ex 27 Sclice 7
	/*
		sl := []int{1, 3, 4, 5, 6, 7, 5}
		fmt.Printf("len: %v | cap: %v\n\r", len(sl), cap(sl))
		fmt.Println("Growing...")
		sl = sl[:0]
		fmt.Printf("len: %v | cap: %v\n\r", len(sl), cap(sl))
		fmt.Println("Growing...")
		sl = sl[:4]
		fmt.Printf("len: %v | cap: %v\n\r", len(sl), cap(sl))
		fmt.Println("Growing...")
		sl = sl[:7]
		fmt.Printf("len: %v | cap: %v\n\r", len(sl), cap(sl))

		fmt.Println("Cut caps...")
		sl = sl[2:]
		fmt.Printf("len: %v | cap: %v\n\r", len(sl), cap(sl))
		fmt.Println("Cut caps...")
		sl = sl[2:]
		fmt.Printf("len: %v | cap: %v\n\r", len(sl), cap(sl))
	*/

	//ex 28 Slice 8 Nil
	/*
		var s []int
		if s == nil {
			fmt.Println("Is nil!")
		}

		var s1 = []int{}
		if s1 == nil {
			fmt.Println("Is nil!")
		} else {
			fmt.Println("Not nil!")
		}
	*/

	//ex 29 Slice 9 make
	/*
		a := make([]int, 5)

		if a == nil {
			fmt.Println("Is nil!")
		} else {
			fmt.Println("Not nil!")
		}

		b := make([]int, 0)

		if b == nil {
			fmt.Println("Is nil!")
		} else {
			fmt.Println("Not nil!")
		}

		c := make([]int, 4, 10)

		if c == nil {
			fmt.Println("Is nil!")
		} else {
			fmt.Println("Not nil!")
		}

		fmt.Println(a, b, c)
	*/

	//ex 30 Slice 10 New
	/*
		sl := new([]int)
		if sl == nil {
			fmt.Println("Is nil!")
		} else {
			fmt.Println("Not nil!")
		}
		fmt.Println(sl)
	*/

	//31 Slice 11 Multy slices
	/*
		brd := [][]string{
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
			[]string{"_", "_", "_"},
		}

		fmt.Println()
		for i := 0; i < len(brd); i++ {
			for j := 0; j < len(brd); j++ {
				fmt.Print(brd[i][j], " ")
			}
			fmt.Println()
		}

		brd[0][0] = "X"
		brd[2][1] = "O"
		brd[0][1] = "X"
		brd[0][2] = "O"
		brd[1][0] = "X"
		brd[2][0] = "O"
		fmt.Println()

		for i := 0; i < len(brd); i++ {
			for j := 0; j < len(brd); j++ {
				fmt.Print(brd[i][j], " ")
			}
			fmt.Println()
		}
	*/

	//ex 32 Slice 12 Append
	/*
		sl := []int{1, 2, 3}
		fmt.Println(sl)
		sl2 := []int{11, 22, 33}
		fmt.Println(sl2)
		sl = append(sl, 4, 5)
		fmt.Println(sl)
		sl = append(sl, sl2...)
		fmt.Println(sl)
	*/

	//ex 33 for
	/*
		sum := 0
		for i := 1; i < 11; i++ {
			sum += i
		}
		fmt.Println(sum)
	*/

	//ex 34 for
	/*
		sum := 1
		for sum < 1000 {
			sum += sum
		}
		fmt.Println(sum)
	*/

	//ex 35 for
	/*
		for {
		}
	*/
	//ex 36 if-else if-else
	/*
		x := 4

		fmt.Scan(&x)

		if x < 3 && x > 0 {
			fmt.Println("1 2")
		} else if x == 100 || == 1000 {
			fmt.Println("100 or 1000")
		} else {
			fmt.Println("Unknown value")
		}
	*/

	//ex 37 if-else if-else
	/*
		hell := 3.1
		if pi := math.Pi; pi < hell {
			fmt.Println("God alive")
		} else {
			fmt.Println("Broke your mind")
		}
		//here pi from if can't be called!!!
	*/

	//ex 38 switch/case
	/*
		fmt.Println("Golang runs on ")
		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X")
		case "linux":
			fmt.Println("GNU/LINUX")
		default:
			fmt.Printf("%s\n\r\n", os)
		}
	*/
}

//ex 12 Functions
/*
func add(a, b int) int {
	return a + b
}
*/

//ex 13 Functions 1
/*
func add(a, b int) int {
	return a + b
}
*/

//ex 14 Functions 2
/*
func swap(x, y string) (string, string) {
	return y, x
}
*/

//ex 15 Functions 3
/*
func swap(x, y string) (string, string) {
	return y, x
}
*/

//ex 16 Functions 4
/*
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
*/

//ex 17 Functions 5 Function as a type
/*
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
*/

//ex 18 Functions 6 Closure
/*
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
*/
