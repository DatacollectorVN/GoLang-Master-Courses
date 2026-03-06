package main

import "fmt"

func main() {
	day := "Sunday"

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("It's the weekend")
	case "Monday":
		fmt.Println("Lots of meetings")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println("--------------------------------")

	age := 20
	switch {
	case age < 18:
		fmt.Println("You are a child")
	case age >= 18 && age < 65:
		fmt.Println("You are a worker")
	default:
		fmt.Println("You are an adult")
	}

	checkType := func(value interface{}) {
		switch v := value.(type) {
		case int:
			fmt.Println("The value is an integer:", v)
		case string:
			fmt.Println("The value is a string:", v)
		case bool:
			fmt.Println("The value is a boolean:", v)
		default:
			fmt.Println("The value is a unknown type:", v)
		}
	}

	checkType(10)
	checkType("Hello")
	checkType(true)
	checkType(10.5)
}
