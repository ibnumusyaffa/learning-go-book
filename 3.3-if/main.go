package main

import "fmt"

func main() {

	//only use this feature to define new variables that are scoped to the if/else statements; anything else would be confusing.
	if n := 7; n == 1 {
		fmt.Println("That's too low", n)
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
	// n is undeclared
	// fmt.Println("n:", n)
}
