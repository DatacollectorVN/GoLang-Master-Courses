package main

import (
	"fmt"
	"time"
)

// nextEmployeeID is used only when creating Employees (see Init and NewEmployee).
// It is package-level so it persists and auto-increments across all new employees.
var nextEmployeeID int = 1

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    float64
	IsActive  bool
	JoinedAt  time.Time
}

// Init sets up the Employee fields, auto-assigns an incremented ID (from nextEmployeeID), and sets JoinedAt.
// Notice the (e *Employee) - this is the pointer receiver.
func (e *Employee) Init(first, last, pos string, salary float64, active bool) *Employee {
	e.ID = nextEmployeeID
	nextEmployeeID++
	e.FirstName = first
	e.LastName = last
	e.Position = pos
	e.Salary = salary
	e.IsActive = active
	e.JoinedAt = time.Now()
	return e
}

// NewEmployee creates a new Employee with an auto-incremented ID. This is the only other place that uses nextEmployeeID.
func NewEmployee(first, last, pos string, salary float64, active bool) *Employee {
	e := &Employee{}
	return e.Init(first, last, pos, salary, active)
}

func (e Employee) GetFullName() string {
	return e.FirstName + " " + e.LastName
}

func main() {
	var nathan, jane Employee
	nathan.Init("Nathan", "Ngo", "Software Engineer", 100000, true)
	jane.Init("Jane", "Doe", "Designer", 80000, true)
	fmt.Printf("Employee nathan: %+v\n", &nathan)
	fmt.Printf("Employee jane: %+v\n", &jane)
}
