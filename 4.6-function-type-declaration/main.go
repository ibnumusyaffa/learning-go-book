package main

import "fmt"

//function type declaration
type opFuncType func(int, int) int

func main() {
	var opMap = map[string]opFuncType{
		"+": add,
	}

	fmt.Println(opMap)

}
func add(i int, j int) int {
	return i + j
}
