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
	fmt.Println(fred.age) // 0, zero value of int

	//same as top
	bob := person{}
	fmt.Println(fred == bob) // true

	bob.name = "bob"
	fmt.Println(bob)

	julia := person{
		"Julia",
		40,
	}

	beth := person{
		age:  30,
		name: "Beth",
	}

	fmt.Println(julia)
	fmt.Println(beth)

	//anonymous struct

	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "hehe"
	person.age = 50
	person.pet = "dog"

	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}

	fmt.Println(pet)

	//asa
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}

	g.name = "Bob"
	g.age = 50

	fmt.Println(f == g) //true
}
