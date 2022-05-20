package main

import "fmt"

func main() {

	func() {
		fmt.Println("integer literal")
		a := 0xF // the hex form (starts with a "0x" or "0X")
		b := 0xF

		c := 017 // the octal form (starts with a "0", "0o" or "0O")
		d := 0o17
		e := 0o17

		f := 0b1111 // the binary form (starts with a "0b" or "0B")
		g := 0b1111

		h := 15 // the decimal form (starts without a "0")
		i := 1500_000

		fmt.Println(a, b, c, d, e, f, g, h, i)
	}()

	func() {
		fmt.Println("floating poiint literal")
		a := 1.23
		b := 01.23 // == 1.23
		c := .23
		d := 1.
		// An "e" or "E" starts the exponent part (10-based).
		e := 1.23e2  // == 123.0
		f := 123e2   // == 12300.0
		g := 123.e+2 // == 12300.0
		h := 1e-1    // == 0.1
		i := .1e0    // == 0.1
		j := 0010e-2 // == 0.1
		k := 0e+5    // == 0.0

		fmt.Println(a, b, c, d, e, f, g, h, i, j, k)
	}()

	func() {
		fmt.Println("rune literal")
		a := 'a' // an English character
		b := 'π'
		c := '众' // a Chinese character

		// 141 is the octal representation of decimal number 97.
		d := '\141'
		// 61 is the hex representation of decimal number 97.
		e := '\x61'
		f := '\u0061'
		g := '\U00000061'

		fmt.Println(a, b, c, d, e, f, g)
	}()

	func() {
		// The interpreted form.
		a := "Hello\nworld!\n\"你好世界\""

		// The raw form.
		b := `Hello
					world!
					"你好世界"`

		fmt.Println(a)
		fmt.Println(b)
	}()

}
