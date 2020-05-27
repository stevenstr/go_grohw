/*
* Author: Stefan
* Date: 05.28.2020 00:10
* Task: Lec5 Structures CW
 */

package main

import (
	"fmt"
)

type user struct {
	firstname string
	lastname  string
	age       int
	contact
}

type contact struct {
	email   string
	zipCode int
}

func main() {
	//ex 1
	/*
		userOne := user{"Stan", "Goose", 12}
		userTwo := user{firstname: "Alladin", lastname: "Alladin"}
		userThrid := user{
			"Garrosh",
			"Hellscream",
			35,
		}

		fmt.Println(userOne)
		fmt.Println(userTwo)

		fmt.Printf("%+vn\r\n", userThrid)
	*/

	//ex 2 zero values
	/*
		var john user
		fmt.Printf("%+v\n\r\n", john) // "" "" 0

		john.firstname = "Jo"
		john.lastname = "Keke"
		fmt.Printf("%+v\n\r\n", john) // fn:"Jo" ln:"Keke" 0
	*/

	//ex 3
	/*
		var john user

		john.firstname = "Jo"
		john.lastname = "Keke"
		john.hisContact.email = "shddhh@gunuil.cum"
		john.hisContact.zipCode = 77700939

		fmt.Printf("%+v\n\r\n", john)
	*/

	//ex 4 Bad readability
	/*
		john := user{
			firstname: "Jon",
			lastname:  "Smith",
			age:       33,
			hisContact: contact{
				email:   "dhjsd@dhfsjdf.dfdf",
				zipCode: 323423423,
			},
		}

		fmt.Printf("%+v\n\r\n", john) //
	*/

	//ex 5 Anonym field
	/*
		cont := contact{
			email:   "ncnncn@nncncn",
			zipCode: 434234234,
		}
		john := user{
			firstname: "Jo",
			lastname:  "Jdldld",
			age:       33,
			contact:   cont,
		}

		fmt.Printf("%+v\n\r\n", john) //
	*/

	//ex 6 Anonym field features
	/*
		cont := contact{
			email:   "ncnncn@nncncn",
			zipCode: 434234234,
		}
		john := user{
			email:     "fffffffffff",
			firstname: "Jo",
			lastname:  "Jdldld",
			age:       33,
			contact:   cont,
		}

		fmt.Printf("%+v\n\r\n", john) //
		fmt.Printf("%+v\n", john.email)
		fmt.Printf("%+v\n", john.contact.email)
	*/

	//ex 7 Anonym field features
	/*
		cont := contact{
			email:   "ncnncn@nncncn",
			zipCode: 434234234,
		}
		john := user{
			//email:     "fffffffffff",
			firstname: "Jo",
			lastname:  "Jdldld",
			age:       33,
			contact:   cont,
		}

		fmt.Printf("%+v\n\r\n", john) //
		fmt.Printf("%+v\n", john.email)
		fmt.Printf("%+v\n", john.contact.email)
	*/

	//ex 8 Methods reciever - Full form
	/*
		cont := contact{
			email:   "ncnncn@nncncn",
			zipCode: 434234234,
		}
		john := user{
			firstname: "Jo",
			lastname:  "Jdldld",
			age:       33,
			contact:   cont,
		}

		fmt.Println("Name should be changed to Kolya because pointer reciever")
		johnPointer := &john
		fmt.Println(johnPointer)
		print(johnPointer)
		johnPointer.updateFirstNameTr("Kolya")
		fmt.Println(johnPointer)
		print(johnPointer)

			//john.show()
			//fmt.Printf("%+v\n", john.email)
			//fmt.Printf("%+v\n", john.contact.email)

		fmt.Println("Still should be kolya because not reciever pointer!")
		john.updateFirstName("Joshua")
		john.show()

	*/

	//ex 9 - Methods Reciever Short form
	cont := contact{
		email:   "ncnncn@nncncn",
		zipCode: 434234234,
	}
	john := user{
		firstname: "Jo",
		lastname:  "Jdldld",
		age:       33,
		contact:   cont,
	}

	john.show()
	john.updateFirstNameTr("Kolya")
	john.show()

}

//show method
func (u user) show() {
	fmt.Printf("%+v\n\r\n", u)
}

//so that you can change value use star(*)
/*
func (u user) updateFirstName(newName string) {
	u.firstname = newName
}
*/

func (u *user) updateFirstNameTr(newName string) {
	(*u).firstname = newName
}
