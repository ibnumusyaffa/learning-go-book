package main

import (
	"errors"
	"fmt"
	"os"
)

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func doThing() (string, error) {
	return "", errors.New("aaaa")
}

func DoSomeThings(val1 int, val2 string) (string, error) {
	_, err := doThing()
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	_, err = doThing()
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	_, err = doThing()
	if err != nil {
		return "", fmt.Errorf("in DoSomeThings: %w", err)
	}
	return "", nil
}
func DoSomeThings2(val1 int, val2 string) (_ string, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomeThings: %w", err)
		}
	}()
	_, err = doThing()
	if err != nil {
		return "", err
	}
	_, err = doThing()
	if err != nil {
		return "", err
	}
	return "", nil
}

func main() {
	err := fileChecker("not_here.txt")
	if err != nil {
		fmt.Println(err) //in fileChecker: open not_here.txt: no such file or directory
		if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
			fmt.Println(wrappedErr) //open not_here.txt: no such file or directory
		}
	}
}
