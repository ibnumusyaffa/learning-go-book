package main

import "fmt"

func main() {
	fmt.Println("shadow variable---------------------------------------")
	a := 10
	b := 10
	if a > 5 {
		fmt.Println("a outer:", a) // 10
		fmt.Println("b outer:", b) // 10
		//shadowing variable
		//create new variable x
		a := 5
		b = 5
		fmt.Println("a inner after redeclare", a) // 5
		fmt.Println("b inner after reassign", b)  // 5
	}

	fmt.Println("a outer", a) //10
	fmt.Println("b outer", b) //5
}
