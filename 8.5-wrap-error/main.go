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
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

// you want to wrap an error with your custom error type, your error type needs to implement the method Unwrap. This method takes in no parameters and returns an error.
func (se StatusErr) Unwrap() error {
	return se.Err
}

func login(uid string, pwd string) error {
	return errors.New("Error")
}

func getData(file string) ([]byte, error) {
	return []byte{1, 2, 3}, errors.New("Error")
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			Err:     err,
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
			Err:     err,
		}
	}
	return data, nil
}

func main() {

}
