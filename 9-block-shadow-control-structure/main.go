package main

import (
	"fmt"
)

func main() {

	fmt.Println("shadow variable---------------------------------------")
	x := 10
	if x > 5 {
		fmt.Println(x) // 10
		//shadowing variable
		//create new variable x
		x := 5         // if x = 5 so outer variable x also change
		fmt.Println(x) // 5
	}
	fmt.Println(x) //10

	fmt.Println("If-------------------------------------------------------")

	if n := 8; n > 1 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}

	fmt.Println("For-------------------------------------------------------")

	// For
	// 1. A complete, C-style for
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	// 2. A condition-only for
	func() {
		i := 1
		for i < 100 {
			fmt.Println(i)
			i = i * 2
		}
	}()

	// 3. An infinite for
	// for {
	// 	fmt.Println("Hello")
	//	if(condition){ break }
	//  if(condition){ continue }
	// }
	//
	// 4. for-range
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		//i index
		//v value
		fmt.Println(i, v)
	}

	//if you dont use index, change to _
	for _, v := range evenVals {
		//i index
		//v value
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k, v := range uniqueNames {
		fmt.Println(k, v)
	}

	//
	fmt.Println("---------------")
	//loop over map is random, so dont rely on order
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	fmt.Println("---------------------")
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			//i = index
			//r = numeric value of letter
			//string(r) = numeric value of letter converted to string
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}

	//Labeling Your for Statements
outer:
	for _, sample := range samples {
	inside:
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}

			if r == 'h' {
				continue inside
			}
		}
		fmt.Println()
	}

	fmt.Println("switch----------------------------------")
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	func() {
		words := []string{"hi", "salutations", "hello"}
		for _, word := range words {
			switch wordLen := len(word); {
			case wordLen < 5:
				fmt.Println(word, "is a short word!")
			case wordLen > 10:
				fmt.Println(word, "is a long word!")
			default:
				fmt.Println(word, "is exactly the right length.")
			}
		}
	}()

}
