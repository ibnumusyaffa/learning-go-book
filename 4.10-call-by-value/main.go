package main

import "fmt"

type person struct {
	age  int
	name string
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

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
	fmt.Println("inside modSlice", s)
}

func main() {
	myPerson := person{}
	myInt := 2
	myString := "Hello"
	modifyFails(myInt, myString, myPerson)

	//struct person does not change,because is passes by value
	fmt.Println(myInt, myString, myPerson)

	fmt.Println("--------------------------------------------------------")
	myMap := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(myMap)
	fmt.Println(myMap)

	fmt.Println("--------------------------------------------------------")

	mySlice := []int{1, 2, 3}
	modSlice(mySlice)
	fmt.Println(mySlice)
}
