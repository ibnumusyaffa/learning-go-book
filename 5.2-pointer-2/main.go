package main

import "fmt"

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func failedUpdate2(p *int) {
	x := 20
	p = &x
}

func update(p *int) {
	*p = 20
}

func main() {
	var f *int // f is nil
	failedUpdate(f)
	fmt.Println(f) // prints nil

	fmt.Println("---------------------")

	x := 10
	failedUpdate2(&x)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)

	fmt.Println("----------------------------------")
	myArr := []int{1, 2, 3, 4}
	fmt.Println(myArr)
	modArr(myArr)
	fmt.Println(myArr)

}

func modArr(arr []int) {
	arr[0] = 99
	arr = append(arr, 10)
	//not affect original array
	fmt.Println("inside mod arr", arr)
}
