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
	fmt.Println("failedUpdate value = ", *value)
}

func update(value *int) {
	*value = 99
}

func updatePerson(p *person) {
	p.LastName = "Musyaffa"
}

func updatePerson2(p person) {
	p.LastName = "updated"
	fmt.Println("inside updatePerson2", p)
}

func main() {
	var a *int // f is nil
	failedUpdate(a)
	fmt.Println("a", a)

	fmt.Println("---------------------------------------------------------")
	b := 10
	failedUpdate(&b)
	fmt.Println("b", b)

	fmt.Println("---------------------------------------------------------")
	c := 10
	update(&c)
	fmt.Println("c", c)
	fmt.Println("----------------------------------------------------------")
	d := person{
		FirstName: "Ibnu",
		LastName:  "M",
	}
	fmt.Println("before d", d)
	updatePerson(&d)
	fmt.Println("after d", d)

	fmt.Println("----------------------------------------------------------")

	e := person{
		FirstName: "Budi",
		LastName:  "M",
	}
	fmt.Println("before e", e)
	updatePerson2(e)
	fmt.Println("after e", e) // not change because is not pointer
}
