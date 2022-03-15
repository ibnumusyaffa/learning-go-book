package main

import (
	"fmt"
)

func main() {

	//Boolean, zero value = false
	var flag bool
	fmt.Println(flag)

	// Integer, zero value = 0
	// int = int64 in 64bit cpu and int32 in 32bit cpu
	// int8 –128 to 127
	// int16 –32768 to 32767
	// int32 –2147483648 to 2147483647
	// int64 –9223372036854775808 to 9223372036854775807
	// uint8 0 to 255
	// uint16 0 to 65536
	// uint32 0 to 4294967295
	// uint64 0 to 18446744073709551615
	// byte is alias of uint8
	var myInteger int
	fmt.Println(myInteger)

	//string,zero value = empty string ''
	var myString string
	fmt.Println(myString)

	//float, zero value = 0.0
	//float32
	//float64
	var myFloat float32
	fmt.Println(myFloat) // 0

}
