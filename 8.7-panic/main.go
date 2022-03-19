package main

import "fmt"

func doPanic(msg string) {
	panic(msg)
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main() {
	// doPanic("hallo")
	for _, val := range []int{1, 2, 0, 6} {
		div60(val)
	}

}
