package main

import "fmt"

func main() {
	studentGrades := map[string]float64{} // Zero value is map[]

	fmt.Println(studentGrades)
	studentGrades = map[string]float64{
		"John": 85.5,
		"Jane": 90.0,
		"Jim":  78.5,
		"Jill": 88.0,
	}
	fmt.Println(studentGrades)
	studentGrades["John"] = 90.0
	fmt.Println(studentGrades)
	studentGrades["John"] = 85.5

	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Println("Alice's grade is", alice)
	} else {
		fmt.Println("Alice's grade is not found")
		fmt.Println("So Alice's grade is", alice)
	}

	config := make(map[string]string) // Zero value is map[]
	fmt.Println(config)

	fmt.Printf("Config: %+v\n %v\n", config, config == nil) // nil means no value / not initialized / points to nothing

	if config == nil {
		fmt.Println("Config is nil")
	} else {
		fmt.Println("Config is not nil")
	}

	var config2 map[string]string // Zero value is nil
	fmt.Println(config2)
	fmt.Printf("Config2: %+v\n %v\n", config2, config2 == nil)
	if config2 == nil {
		fmt.Println("Config2 is nil")
	} else {
		fmt.Println("Config2 is not nil")
	}

	var config3 = map[string]string{} // Zero value is map[]
	fmt.Println(config3)
	fmt.Printf("config3: %+v\n %v\n", config3, config3 == nil)
	if config3 == nil {
		fmt.Println("config3 is nil")
	} else {
		fmt.Println("config3 is not nil")
	}
	fmt.Println("--------------------------------")

	config["port"] = "8080"
	config["host"] = "localhost"
	config["protocol"] = "http"
	fmt.Println(config)
	delete(config, "protocol")
	fmt.Println(config)
	fmt.Println(config["port"])
	fmt.Println(config["host"])
	fmt.Println(config["protocol"])
}

/*
Three Ways to Declare a MapTo keep it clear, here is how the three different declarations look to the Go runtime:
Syntax					Status			Memory Allocated?	Result of == nil
var m map[string]string	Nil Map			Not allocated		true
m := make(map[string]string)	Empty Map		Allocated			false
m := map[string]string{}		Empty Map		Allocated			false
*/
