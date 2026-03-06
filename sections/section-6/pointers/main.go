package main

import "fmt"

func main() {
	age := 20
	agePtr := &age
	fmt.Println(agePtr)
	fmt.Println(*agePtr) // dereference the pointer to get the value

	fmt.Println("--------------------------------")
	*agePtr = 30 // change the value of the variable through the pointer
	fmt.Println(age)
	fmt.Println(*agePtr)
}
