package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d", c.total)
}

func main() {
	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}
	output := p.String()
	fmt.Println(output)

	fmt.Println("--------------------------------------------------------------------")

	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
	fmt.Println("--------------------------------------------------------------------")
	var d Counter
	doUpdateWrong(d)
	doUpdateRight(&d)
	fmt.Println(d.String())
}

func doUpdateWrong(c Counter) {
	//it just copy, does not reflect variable outside function
	c.Increment()
	// fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	// fmt.Println("in doUpdateRight:", c.String())
}
