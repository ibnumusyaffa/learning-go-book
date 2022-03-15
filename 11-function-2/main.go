package main

import (
	"fmt"
	"sort"
)

func main() {

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	fmt.Println("------------------------------------------")

	twoBase := makeMult(2)
	threeBase := makeMult(3)

	fmt.Println(twoBase(5), threeBase(3)) // 5 x 2, 3 x 3

}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
