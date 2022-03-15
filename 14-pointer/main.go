package main

import "fmt"

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

type animal struct {
	FirstName  string
	MiddleName int
	LastName   int
}

func main() {

	var x int = 123
	//get pointer of x
	pointerX := &x
	fmt.Println(pointerX)  // prints a memory address
	fmt.Println(*pointerX) // prints 10

	fmt.Println("-----------------------------------------------")
	var z string = "aaa"
	//pointer type of string
	var pointerZ *string
	//The zero value for a pointer is nil
	fmt.Println(pointerZ)
	pointerZ = &z
	fmt.Println(pointerZ)
	fmt.Println(*pointerZ) // deferencing will panic when pointer is nill
	fmt.Println("?????-----------------------------------------------")
	var a = new(int)
	fmt.Println(a == nil) // prints false
	fmt.Println(*a)       // 0
	fmt.Println("-----------------------------------------------")
	// Before dereferencing a pointer, you must make sure that the pointer is non-nil. Your program will panic if you attempt to dereference a nil pointer:
	// var x *int
	// fmt.Println(x == nil) // prints true
	// fmt.Println(*x)       // panics
	p := animal{FirstName: "dog"}

	mutStruct(p)
	fmt.Println(p)
	// fmt.Println(*p)

	fmt.Println("------------------------------------------------")
	func() {
		x := 10
		y := x
		y = 20

		fmt.Println(x, y)
	}()

	fmt.Println("------------------------------------------")
	y := 10
	myFun(&y)
	fmt.Println("after", y)

}

func myFun(myInt *int) {
	fmt.Println("before", *myInt)
	*myInt = 20
	// fmt.Println(*myInt)
}

func mutStruct(a animal) {
	a.FirstName = "new name"
}
