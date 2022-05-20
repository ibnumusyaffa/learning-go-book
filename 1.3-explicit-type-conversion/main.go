package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Most languages that have multiple numeric types automatically convert from one to another when needed. This is called automatic type promotion, and while it seems very convenient, it turns out that the rules to properly convert one type to another can get complicated and produce unexpected results. As a language that values clarity of intent and readability, Go doesn’t allow automatic type promotion between variables. You must use a type conversion when variable types do not match. Even different-sized integers and floats must be converted to the same type to interact. This makes it clear exactly what type you want without having to memorize any type conversion rules

	var myInt int = 10
	var myFloat float64 = 30.2
	//convert x to float64
	var myFloat2 float64 = float64(myInt) + myFloat
	//convert y to int
	var myInt2 int = myInt + int(myFloat2)

	toString := strconv.Itoa(123)

	fmt.Println("myInt:", myInt)
	fmt.Println("myFloat:", myFloat)
	fmt.Println("myFloat2:", myFloat2)
	fmt.Println("myInt2", myInt2)
	fmt.Println(toString)

	// This strictness around types has other implications. Since all type conversions in Go are explicit, you cannot treat another Go type as a boolean. In many languages, a nonzero number or a nonempty string can be interpreted as a boolean true. Just like automatic type promotion, the rules for “truthy” values vary from language to language and can be confusing. Unsurprisingly, Go doesn’t allow truthiness. In fact, no other type can be converted to a bool, implicitly or explicitly. If you want to convert from another data type to boolean, you must use one of the comparison operators (==, !=, >, <, <=, or >=). For example, to check if variable x is equal to 0, the code would be x == 0. If you want to check if string s is empty, use s == "".
	if myFloat > 0 {
		fmt.Println("d is greater than 0", myFloat)
	}

}
