package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
}

func failedUpdate(value *int) {
	x := 99
	value = &x
	//pointer is update, but not reflect in outside function
	// fmt.Println("failedUpdate value = ", *value)
}

func update(value *int) {
	*value = 99
}

func updatePerson(p *person) {
	p.LastName = "Musyaffa"
}

func main() {
	var a *int // f is nil
	failedUpdate(a)
	fmt.Println("a", a)

	b := 10
	failedUpdate(&b)
	fmt.Println("b", b)

	c := 10
	update(&c)
	fmt.Println("c", c)

	d := person{
		FirstName: "Ibnu",
		LastName:  "M",
	}
	updatePerson(&d)
	fmt.Println("d", d)
}
