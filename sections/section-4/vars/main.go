package main

import "fmt"

func main() {
	var greeting string // zero value is ""
	fmt.Println(greeting)
	greeting = "Hello World!"
	fmt.Println(greeting)

	var count int // zero value is 0
	fmt.Println(count)
	count = 10
	fmt.Println(count)

	var isActive bool // zero value is false
	fmt.Println(isActive)
	isActive = true
	fmt.Println(isActive)

	var price float64 // zero value is 0.0
	fmt.Println(price)
	price = 10.5
	fmt.Println(price)

	email := "datacollectoriu@gmail.com" // short declaration operator
	fmt.Println(email)

	var year = 2026
	fmt.Println(year)

	age, name := 26, "Nathan"
	fmt.Println(age, name)

	var (
		firstName = "Nathan"
		lastName  = "Ngo"
		age2      = 26
		isActive2 = false
	)
	fmt.Println(firstName, lastName, age2, isActive2)
}
