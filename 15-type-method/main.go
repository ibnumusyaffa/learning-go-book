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

//methods on user-defined types
func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

func (p *Person) changeName() {
	p.FirstName = "new"
}

type Score int
type Converter func(string) Score
type TeamScores map[string]Score

func (s Score) my() {
	fmt.Println(s)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

func main() {

	p := Person{
		FirstName: "Fred",
		LastName:  "Fredson",
		Age:       52,
	}
	output := p.String()
	fmt.Println(output)

	p.changeName()
	fmt.Println(p)

	fmt.Println("------------------------------------------")
	var c Counter
	c.Increment()
	c.Increment()
	fmt.Println(c.total)

	var b Counter
	doUpdateWrong(b)
	fmt.Println(b.total)

	doUpdateRight(&b)
	fmt.Println(b.total)

	fmt.Println("------------------------------------------")

	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
	fmt.Println(it)

	fmt.Println("---------------")

	var myScore Score = 10
	myScore.my()
}

func doUpdateWrong(c Counter) {
	c.Increment()
}

func doUpdateRight(c *Counter) {
	c.Increment()
}
