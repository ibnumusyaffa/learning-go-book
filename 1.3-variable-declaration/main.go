package main

import "fmt"

func main() {
	var a int = 10
	var b = 10
	var c int
	d := 10

	fmt.Println(a, b, c, d)

	var f, g int
	fmt.Println(f, g)

	var h, i, j = true, false, "no!"
	fmt.Println(h, i, j)

	k, l, m := true, false, "no!"
	fmt.Println(k, l, m)

	fmt.Println("------------------------------------------------------------")
	const n = 10
	const o int = 10

	// const cannot be array, slice, map, or struct
	const (
		idKey   = "id"
		nameKey = "name"
	)

	// Untyped constant
	const p = 10
	// All the following assignments are legal
	var q int = p
	var r float64 = p
	var s byte = p
	fmt.Println(p, q, r, s)

	// typed constant
	const t int = 10
	// error, cannot assign t (int) to variable u (float64)
	// var u float64 = t
}
