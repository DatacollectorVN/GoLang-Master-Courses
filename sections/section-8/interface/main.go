package main

import "fmt"

type Person interface {
	GetName() string
	GetID() int
}

type Employee struct {
	ID   int
	Name string
}

func (e Employee) GetName() string {
	return e.Name
}

func (e Employee) GetID() int {
	return e.ID
}

type BusinessPerson struct {
	ID   int
	Name string
}

func (b BusinessPerson) GetName() string {
	return b.Name
}

func (b BusinessPerson) GetID() int {
	return b.ID
}

func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

func main() {
	employee := Employee{ID: 1, Name: "John Doe"}
	businessPerson := BusinessPerson{ID: 2, Name: "Jane Doe"}
	displayPerson(employee)
	displayPerson(businessPerson)
	fmt.Println(employee.GetID())
	fmt.Println(businessPerson.GetID())
}
