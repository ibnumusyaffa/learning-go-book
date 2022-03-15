package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	func() {
		result, err := division(10, 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(result)
	}()

	func() {
		//ignore value
		_, err := division(10, 0)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}()

}

func division(numerator int, denominator int) (int, error) {
	if denominator == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return numerator / denominator, nil
}

// named return values
// When you supply names to your return values, what you are doing is pre-declaring variables that you use within the function to hold the return values. They are written as a comma-separated list within parentheses. You must surround named return values with parentheses, even if there is only a single return value. Named return values are initialized to their zero values when created. This means that we can return them before any explicit use or assignment.
func division2(numerator int, denominator int) (result int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return result, err
	}

	result = numerator / denominator
	err = nil
	return result, err
}
