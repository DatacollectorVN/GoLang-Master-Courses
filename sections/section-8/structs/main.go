package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    float64
	IsActive  bool
	JoinedAt  time.Time
}

func newEmployee(id int, firstName, lastName, position string, salary float64, isActive bool, joinedAt time.Time) *Employee {
	return &Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  position,
		Salary:    salary,
		IsActive:  isActive,
		JoinedAt:  joinedAt,
	}
}

func main() {
	nathan := Employee{
		ID:        1,
		FirstName: "Nathan",
		LastName:  "Ngo",
		Position:  "Software Engineer",
		Salary:    100000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	fmt.Printf("Nathan: %+v\n", nathan)
	fmt.Println(nathan.FirstName)
	fmt.Println(nathan.LastName)
	fmt.Println(nathan.Position)
	fmt.Println(nathan.Salary)
	fmt.Println(nathan.IsActive)
	fmt.Println(nathan.JoinedAt)

	nhu := newEmployee(2, "Nhu", "Nguyen", "Doctor", 100000, true, time.Now())
	fmt.Printf("Nhu: %+v\n", nhu) // return a pointer to the employee
	fmt.Println(nhu.FirstName)
	fmt.Println(nhu.LastName)
	fmt.Println(nhu.Position)
	fmt.Println(nhu.Salary)
	fmt.Println(nhu.IsActive)
	fmt.Println(nhu.JoinedAt)

	nathan_2 := nathan
	fmt.Printf("Nathan_2: %+v\n", nathan_2)
	nathan_2.IsActive = false
	fmt.Println(nathan_2.IsActive)
	fmt.Println(nathan.IsActive)

	fmt.Println("--------------------------------")
	nathanPtr := &nathan
	fmt.Printf("NathanPtr: %+v\n", *nathanPtr)
	nathanPtr.IsActive = false
	fmt.Printf("NathanPtr: %+v\n", *nathanPtr)
	fmt.Println(nathan.IsActive)
}
