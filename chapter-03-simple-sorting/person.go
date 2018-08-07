package main

import "fmt"

type Person struct {
	lastName  string
	firstName string
	age       int
}

func NewPerson(last, first string, age int) *Person {
	p := Person{lastName: last, firstName: first, age: age}
	return &p
}

func (p *Person) Display() {
	fmt.Printf("    Last Name: %s, First Name: %s, Age: %v", p.lastName, p.firstName, p.age)
}
