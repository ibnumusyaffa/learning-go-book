package main

import (
	"fmt"
	"io"
)

type MyInt int

func main() {
	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}

	fmt.Println(i)

	fmt.Println("------------------------------------------------------------")

	var b interface{}
	var mine MyInt = 20
	b = mine
	b2, ok := b.(MyInt)
	//ok = false if b cannot convert to myInt
	if !ok {
		return
	}
	// fmt.Println(b + 1) error
	fmt.Println(b2 + 1)

}

func doThings(i interface{}) {
	switch j := i.(type) {
	case nil:
		// i is nil, type of j is interface{}
		fmt.Println(j)
	case int:
		// j is of type int
	case MyInt:
		// j is of type MyInt
	case io.Reader:
		// j is of type io.Reader
	case string:
		// j is a string
	case bool, rune:
		// i is either a bool or rune, so j is of type interface{}
	default:
		// no idea what i is, so j is of type interface{}
	}
}
