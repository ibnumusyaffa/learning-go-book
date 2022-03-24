package main

import (
	"errors"
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func testErr() error {
	return errors.New("Error")
}

func LoginAndGetData() (string, error) {
	err := testErr()
	if err != nil {
		return "", StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", "s"),
		}
	}
	err = testErr()
	if err != nil {
		return "", StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", "s"),
		}
	}
	return "done", nil
}
func main() {

}
