package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	//Passing Functions as Parameters
	people := []Person{
		{"Cat", "Patterson", 37},
		{"Aracy", "Bobbert", 23},
		{"Bred", "Fredson", 18},
	}
	fmt.Println(people)
	sort.Slice(people, func(i int, j int) bool {
		return people[i].FirstName < people[j].FirstName
	})
	fmt.Println(people)

	twoBase := makeMult(2)
	threeBase := makeMult(3)

	fmt.Println(twoBase(10))   // 2 * 3
	fmt.Println(threeBase(10)) // 2 * 10

}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
