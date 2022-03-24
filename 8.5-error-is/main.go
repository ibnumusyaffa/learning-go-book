package main

import (
	"errors"
	"fmt"
)

var ErrBadInput = errors.New("bad input")

func validateInput(input string) error {
	if input == "abc" {
		return fmt.Errorf("validateInput: %w", ErrBadInput)
	}
	return nil
}

func main() {
	input := "abc"

	err := validateInput(input)
	if errors.Is(err, ErrBadInput) {
		fmt.Println("bad input error")
	}
}
