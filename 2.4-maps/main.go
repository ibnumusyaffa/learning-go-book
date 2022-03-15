package main

import (
	"fmt"
)

func main() {

	func() {
		var nilMap map[string]int
		fmt.Println(nilMap == nil)     //true
		fmt.Println(nilMap["noexits"]) // 0
		//panic: assignment to entry in nil map
		// nilMap["t"] = 1
		// fmt.Println(nilMap)
	}()
	fmt.Println("------------------------------------------------------")
	func() {
		totalWins := map[string]int{}
		fmt.Println(totalWins == nil) //false
		fmt.Println(totalWins)        // map[]
		//not panic, because is not nill
		totalWins["aaa"] = 100
		fmt.Println(totalWins)
		fmt.Println(totalWins["noexits"])

	}()

	fmt.Println("---------------create maps with make------------------")
	func() {
		ages := make(map[int][]string, 10)

		fmt.Println(ages)
		fmt.Println(ages == nil) //false
		fmt.Println(len(ages))   // stil 0
	}()

	fmt.Println("--------------------reading and writin maps------------------")
	func() {
		total := map[string]int{}
		total["anton"] = 1
		total["budi"] = 2
		fmt.Println("before:", total)
		total["anton"]++
		total["celia"] = 3
		fmt.Println("after:", total)

		//get value
		v, ok := total["world"]
		fmt.Println(v, ok)

		//delete from map
		//If the key isn’t present in the map or if the map is nil, nothing happens. The delete function doesn’t return a value.
		delete(total, "anton")
		fmt.Println(total)
	}()

	fmt.Println("----------------using maps as sets------------------")
	func() {
		intSet := map[int]bool{}
		vals := []int{1, 2, 3, 4, 5, 6}
		for _, v := range vals {
			intSet[v] = true
		}

		fmt.Println("vals:", vals)
		fmt.Println("intSet:", intSet)
		fmt.Println(len(vals), len(intSet))

		if intSet[3] {
			fmt.Println("3 is in the set")
		}

	}()

}
