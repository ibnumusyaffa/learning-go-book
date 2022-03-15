package main

import (
	"fmt"
)

func main() {

	var a [3]int
	fmt.Println(a) //zero value for all element = [0 0 0]

	var b = [3]int{10, 20, 30}
	fmt.Println(b)

	//you can use ... instead
	var c = [...]int{1, 2, 3}
	var d = [3]int{1, 2, 3}
	//true jika isi elemen semua sama
	//error jika ukuran array beda
	fmt.Println(c == d) // prints true

	var e = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(e) // [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15].

	//Go only has one-dimensional arrays, but you can simulate multidimensional arrays:
	var f [2][3]int
	fmt.Println(f) // zero value ,[[0 0 0] [0 0 0]]
	f[0] = [3]int{1, 2, 3}
	f[0][0] = 7
	fmt.Println(f[0])
	fmt.Println(f[0][0])

	//This declares f to be an array of length 2 whose type is an array of ints of length 3. This sounds pedantic, but there are languages with true matrix support; Go isn’t one of them.
	fmt.Println(len(a))
}
