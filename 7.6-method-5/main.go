package main

import (
	"fmt"
)

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

// func (m Manager) FindNewEmployees() []Employee {
// 	// do business logic
// }

func main() {

	// Note that Manager contains a field of type Employee, but no name is assigned to that field. This makes Employee an embedded field. Any fields or methods declared on an embedded field are promoted to the containing struct and can be invoked directly on it
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)            // prints 12345
	fmt.Println(m.Employee.ID)   // prints 12345, if has sampe property to outer struct, must access with full path
	fmt.Println(m.Description()) // prints Bob Bobson (12345)
	fmt.Println(m.Reports)
}
