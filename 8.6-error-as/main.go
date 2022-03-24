package main

import (
	"errors"
	"fmt"
)

type BadInputError struct {
	input string
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("bad input: %s", e.input)
}

func validateInput(input string) error {
	if input == "abc" {
		return fmt.Errorf("validateInput: %w", &BadInputError{input: input})
	}
	return nil
}

func main() {
	input := "abc"

	err := validateInput(input)
	var badInputErr *BadInputError
	if errors.As(err, &badInputErr) {
		fmt.Printf("bad input error occured: %s\n", badInputErr)
	}
}
