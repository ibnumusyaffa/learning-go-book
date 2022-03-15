package main

import "fmt"

func main() {
	//fmt.Println(n),there is no n variable

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

	//What we are seeing is special behavior from iterating over a string with a for-range loop. It iterates over the runes, not the bytes
	sample := "apple_π!"
	for i, r := range sample {
		fmt.Println(i, string(r))
	}

	fmt.Println("")

	fmt.Println("Break and continue-------------------------------------------------------")

	counter := 1
	for {
		if counter == 1 {
			//exit loop
			break
		}
	}

	for i := 1; i <= 3; i++ {
		if i == 2 {
			fmt.Println("skips directly to the next iteration")
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("Labelling statement--------------------------------------------------------------")

	//Labeling Your for Statements
	samples := []string{"hello", "apple_π!"}
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
}
