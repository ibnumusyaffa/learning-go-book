package main

import (
	"fmt"
)

func main() {

	var a = []int{10, 20, 30}
	fmt.Println("a:", a)

	var b = []int{2, 2, 3, 3, 5: 7}

	fmt.Println("b:", b)

	b[0] = 10
	fmt.Println(b[0])

	var multiArray [][]int
	multiArray = append(multiArray, []int{1, 2, 3})
	multiArray = append(multiArray, []int{1, 2, 3})
	fmt.Println("multi arr", multiArray)

	//zero value = nil
	var c []int
	fmt.Println(c == nil) //true

	//array with len 0 []
	var d = []int{}
	fmt.Println(d == nil) // false
	fmt.Println(len(d))
	fmt.Println(cap(d))

	//compile time error
	// var e1 = []int{1, 2}
	// var e2 = []int{1, 2}
	// fmt.Println(e1 == e2)

	//append
	var f []int
	f = append(f, 1)
	f = append(f, 2, 3)
	fmt.Println(f)
	g := []int{3, 4}
	f = append(f, g...)
	fmt.Println(f)

	//create slice with make
	h := make([]int, 0)
	fmt.Println(h)        // []
	fmt.Println(h == nil) //false

	//create slice with length 5, capacity 10
	i := make([]int, 5, 10) // [0,0,0,0,0]

	fmt.Println("my init", i)
	fmt.Println(len(i))
	fmt.Println(cap(i))

	//panic index out of range
	// fmt.Println(i[10])

	fmt.Println("-----------------------------------")
	func() {
		x := []int{1, 2, 3, 4}
		y := x[0:1]
		fmt.Println("y:", y)
		//x[0] juga berubah
		y[0] = 1000
		fmt.Println("x:", x)
		fmt.Println("y:", y)

	}()

	fmt.Println("-----------------------------------")
	func() {
		x := []int{1, 2, 3, 4}
		y := x[0:1]
		//x[0] juga berubah
		y[0] = 1000
		y = append(y, 2000)
		fmt.Println("x:", x)
		fmt.Println("y:", y)

	}()

	fmt.Println("-----------------------------------")
	func() {
		x := []int{1, 2, 3, 4}
		y := x[0:1:1]
		//x[0] juga berubah
		// y[0] = 1000
		y = append(y, 2000)
		y = append(y, 3000)
		fmt.Println("x:", x)
		fmt.Println("y:", y)

	}()

	fmt.Println("convert array to slice-----------------------------------")
	func() {
		x := [4]int{5, 6, 7, 8}
		y := x[:1]
		x[0] = 88

		y = append(y, 10000)
		y = append(y, 20000)
		fmt.Println("x:", x)
		fmt.Println("y:", y)
	}()
	fmt.Println("copy-----------------------------------")
	func() {
		src := []int{1, 2, 3, 4, 5, 6}

		dest1 := make([]int, 3)
		fmt.Println(dest1)
		var dest2 []int
		fmt.Println(dest2)

		//only return the number of element has copied
		num1 := copy(dest1, src)
		fmt.Println(dest1, num1)

		num2 := copy(dest2, src)
		fmt.Println(dest2, num2)

		x := []int{1, 2, 3, 4}
		y := make([]int, 2)
		copy(y, x[2:])
		fmt.Println(y)

		// x := []int{1, 2, 3, 4}
		// num = copy(x[:3], x[1:])
		// fmt.Println(x, num) [2 3 4 4] 3.
	}()
}
