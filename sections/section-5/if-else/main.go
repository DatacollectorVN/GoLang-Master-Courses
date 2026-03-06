package main

import "fmt"

func main() {
	age := 20
	if age < 18 {
		fmt.Println("You are a child")
	} else if age >= 18 && age < 65 {
		fmt.Println("You are a worker")
	} else {
		fmt.Println("You are an adult")
	}

	fmt.Println("--------------------------------")

	userAcess := map[string]bool{
		"admin": true,
		"user":  false,
	}
	// ok is a boolean that is true if the key exists in the map, and false if it doesn’t.
	if hasAccess, ok := userAcess["admin"]; ok && hasAccess {
		fmt.Println("You are an admin")
	} else {
		fmt.Println("You are not an admin")
	}
	// can not print the variables here because they are scoped to the if statement
	// fmt.Printf("hasAccess: %t, ok: %t\n", hasAccess, ok)
}
