package main

import "fmt"

type foo struct {
	FirstName string
}

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func main() {
	var a int32 = 10
	var b bool = true
	pointerA := &a
	pointerB := &b
	var pointerC *string
	fmt.Println("pointerA:", pointerA)
	fmt.Println("pointerB:", pointerB)
	fmt.Println("pointerC:", pointerC)

	fmt.Println("PointerA value:", *pointerA)

	if pointerC != nil {
		//*pointerC will panic because *pointerC is nil
		fmt.Println("PointerC Value:", *pointerC)
	}

	//The zero value for a pointer is nil. We’ve seen nil a few times before, as the zero value for slices, maps, and functions. All of these types are implemented with pointers. (Two more types, channels and interfaces, are also implemented with pointers

	fmt.Println("-----------------------------------------------------------------------------")

	var d = new(int)
	fmt.Println("d == nill :", d == nil) // prints false
	fmt.Println("*d value", *d)          // prints 0

	e := &foo{}
	fmt.Println("e:", e)

	var f string
	g := &f
	fmt.Println("g:", g)

	// p := person{
	// 	FirstName:  "Pat",
	// 	MiddleName: "Perry", // This line won't compile
	// 	LastName:   "Peterson",
	// }

	h := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"), // This works
		LastName:   "Peterson",
	}

	fmt.Println("h", h)
}

func stringp(s string) *string {
	return &s
}
