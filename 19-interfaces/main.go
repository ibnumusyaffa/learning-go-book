package main

import (
	"fmt"
)

// Interfaces are usually named with “er” endings. We’ve already seen fmt.Stringer, but there are many more, including io.Reader, io.Closer, io.ReadCloser, json.Marshaler, and http.Handler.
type Stringer interface {
	String() string
}

type LogicProvider struct {
}

func (lp LogicProvider) Process(data string) string {
	return "aaa"
}

type Logic interface {
	Process(string) string
}

//----------------------------
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data from somewhere
	data := ""
	c.L.Process(data)
}

func main() {

	c := Client{
		L: LogicProvider{},
	}

	c.Program()

	//

	//empty interface
	var i interface{}
	i = 20
	i = "hello"
	i = struct {
		FirstName string
		LastName  string
	}{"Fred", "Fredson"}

	fmt.Println(i)

	fmt.Println("---------------------------------")
	var k interface{} = 42

	//type assertion, is ok idiom to prefent panic
	v, ok := k.(string)
	//panic if not string
	//v := k.(string)

	if !ok {
		fmt.Println("Wrong type assertion!")
	} else {
		fmt.Println(v)
	}
}

// When an interface could be one of multiple possible types, use a type switch instead:
// func doThings(i interface{}) {
// 	switch j := i.(type) {
// 	case nil:
// 			// i is nil, type of j is interface{}
// 	case int:
// 			// j is of type int
// 	case MyInt:
// 			// j is of type MyInt
// 	case io.Reader:
// 			// j is of type io.Reader
// 	case string:
// 			// j is a string
// 	case bool, rune:
// 			// i is either a bool or rune, so j is of type interface{}
// 	default:
// 			// no idea what i is, so j is of type interface{}
// 	}
// }
