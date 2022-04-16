package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	ordId      int
	customerId int
}

func main() {
	// var x int
	// xt := reflect.TypeOf(x) // reflect.Type instance
	// fmt.Println(xt.Name())  // returns int ,
	// f := Foo{}
	// ft := reflect.TypeOf(f)
	// fmt.Println(ft.Name()) // returns Foo
	// xpt := reflect.TypeOf(&x)
	// fmt.Println(xpt.Name()) // returns an empty string

	// var x int
	// xpt := reflect.TypeOf(&x)      // reflect.Type instance
	// fmt.Println(xpt.Name())        // returns an empty string
	// fmt.Println(xpt.Kind())        // returns reflect.Ptr
	// fmt.Println(xpt.Elem().Name()) // returns "int"
	// fmt.Println(xpt.Elem().Kind()) // returns reflect.Int

	type Foo struct {
		A int    `myTag:"value"`
		B string `myTag:"value2"`
	}

	var f Foo
	ft := reflect.TypeOf(f)
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(),
			curField.Tag.Get("myTag"))
	}

}
