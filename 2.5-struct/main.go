package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {

	var fred person
	fmt.Println("fred", fred) // 0, zero value of int

	//same as top
	bob := person{}
	fmt.Println("bob", bob)

	//When using this struct literal format, a value for every field in the struct must be specified, and the values are assigned to the fields in the order they were declared in the struct definition
	julia := person{
		"Julia",
		40,
	}

	//The second struct literal style looks like the map literal style:
	beth := person{
		age:  30,
		name: "Beth",
	}

	//read and write struct
	julia.name = "Bob"
	fmt.Println(beth.name)

	fmt.Println("------------------------anonymous struct--------------------------")
	//anonymous struct
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "andy"
	person.age = 50
	person.pet = "dog"

	fmt.Println("person", person)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(pet)

	fmt.Println("------------------------comparing--------------------------")

	func() {

		type person struct {
			name string
			age  int
		}

		var fred person
		fmt.Println("fred", fred)
		//same as top
		bob := person{}
		fmt.Println("bob", bob)

		fmt.Println("fred == bob :", fred == bob) // true

		bob.name = "bob"
		fmt.Println("bob", bob)
		fmt.Println("fred == bob :", fred == bob) // false

		bob.name = ""
		fmt.Println("bob", bob)
		fmt.Println("fred == bob :", fred == bob) // true
	}()

	fmt.Println("-------------------------converting--------------------------")

	func() {
		//Just like Go doesn’t allow comparisons between variables of different primitive types, Go doesn’t allow comparisons between variables that represent structs of different types. Go does allow you to perform a type conversion from one struct type to another if the fields of both structs have the same names, order, and types.

		//We can use a type conversion to convert an instance of firstPerson to secondPerson, but we can’t use == to compare an instance of firstPerson and an instance of secondPerson, because they are different types:
		type firstPerson struct {
			name string
			age  int
		}
		type secondPerson struct {
			name string
			age  int
		}

		//We can’t convert an instance of firstPerson to thirdPerson, because the fields are in a different order:
		type thirdPerson struct {
			age  int
			name string
		}

		//We can’t convert an instance of firstPerson to fourthPerson because the field names don’t match:
		type fourthPerson struct {
			firstName string
			age       int
		}
		//Finally, we can’t convert an instance of firstPerson to fifthPerson because there’s an additional
		type fifthPerson struct {
			name          string
			age           int
			favoriteColor string
		}

		antony := firstPerson{
			age:  30,
			name: "antony",
		}

		beth := secondPerson{
			age:  30,
			name: "Beth",
		}

		celia := thirdPerson{
			age:  30,
			name: "celia",
		}

		fmt.Println(beth == secondPerson(antony)) //false

		fmt.Println(celia)
		// fmt.Println(celia == antony) //error

	}()

}
