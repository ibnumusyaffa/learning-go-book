package main

import "fmt"

func main() {

	var a string = "Hello there"
	var b byte = a[6]
	//var c string = a[0] //error
	// var c string = string(a[0]) //not error
	var c string = a[0:5]
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	var d string = "Hello 😃"
	// This code prints out 10, not 7, because it takes four bytes to represent the sun with smiling face emoji in UTF-8.
	fmt.Println("len:", len(d))

	fmt.Println("---------------------------------------------")
	func() {
		var a rune = 'x'
		var s string = string(a)
		var b byte = 'y'
		var s2 string = string(b)

		fmt.Println("a:", a)
		fmt.Println("s:", s)
		fmt.Println("b:", b)
		fmt.Println("s2:", s2)
	}()

	fmt.Println("---------------------------------------------")
	func() {
		var x int = 65
		var y = string(x)
		fmt.Println(y) //A not "65"
	}()
}
