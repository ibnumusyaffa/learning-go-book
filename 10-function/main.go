package main

import (
	"errors"
	"fmt"
)

//function type declaration
type opFuncType func(int, int) int

func add(i int, j int) int { return i + j }
func main() {

	fmt.Println(variadicFunction(1, 2, 3, 4, 5))

	fmt.Println(divAndRemainder3(10, 2))

	fmt.Println("---------------------")
	//inline type declaration
	var opMap1 = map[string]func(int, int) int{
		"+": add,
	}
	fmt.Println(opMap1)

	var opMap2 = map[string]opFuncType{
		"+": func(i int, j int) int {
			return i + j
		},
	}
	fmt.Println(opMap2)
	opFunc, ok := opMap2["+"]
	if !ok {
		fmt.Println("unsupported operator:")
	}
	result := opFunc(1, 2)
	fmt.Println("result", result)

	//anonymous function
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}
}

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

// When you have multiple input parameters of the same type,
// you can write your input parameters like this
func div2(numerator, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

func variadicFunction(vals ...int) []int {
	return vals
}

//multiple return values
func divAndRemainder(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

//named return values
func divAndRemainder2(numerator int, denominator int) (result int, remainder int,
	err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}

func divAndRemainder3(numerator, denominator int) (result int, remainder int,
	err error) {
	// assign some values
	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}
