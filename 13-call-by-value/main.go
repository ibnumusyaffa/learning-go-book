package main

import (
	"fmt"
)

// You might hear people say that Go is a call by value language and wonder what that means.
// It means that when you supply a variable for a parameter to a function, Go always makes a copy of the value of the variable (except maps and slices)

type person struct {
	age  int
	name string
}

func main() {

	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	fmt.Println("---------------------------")
	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	k := []int{1, 2, 3}
	k = modSlice(k)
	fmt.Println(k)
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) []int {
	for k, v := range s {
		s[k] = v * 2
	}
	//not work if not return
	s = append(s, 10)

	return s
}
