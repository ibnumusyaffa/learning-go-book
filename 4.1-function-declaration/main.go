package main

import "fmt"

func main() {
	result := div(5, 2)
	fmt.Println(result)
}

//When you have multiple input parameters of the same type, you can write your input parameters like this:
//func div(numerator, denominator int) int {
func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}
