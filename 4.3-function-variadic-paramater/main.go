package main

import "fmt"

func main() {
	variadic(1, 2, 3, 4)
}

func variadic(vals ...int) {
	fmt.Println(vals)
	fmt.Printf("vals: %T\n", vals) //[] int, slice of int
}
