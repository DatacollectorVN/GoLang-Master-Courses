package main

import "fmt"

type Number interface {
	int | float64 | float32 | string
}

func Sum[T Number](numbers ...T) T {
	var total T
	for _, number := range numbers {
		var result any
		/*
			any(total)
				Converts total to any (same as interface{}). That turns it into an interface value that still holds the
				same underlying value and its runtime type, but the compiler no longer treats it as T.
				So you can inspect the real type at runtime.

			.(type)
				This is special syntax: it’s only allowed in a type switch. It doesn’t return a value;
				it’s the “what type is it?” expression the switch branches on.
		*/
		switch v := any(total).(type) {
		case string:
			n := any(number).(string)
			if v == "" {
				result = n
			} else {
				result = v + " " + n
			}
		case int:
			/*
				any(total)
					Wraps total in an interface value of type any.
					The value inside doesn’t change; you’re just viewing it as any.
				.(int)
					This is a type assertion. It says: “This any value is holding an int; give me that int.”
					It does not convert (e.g. it does not turn the string "42" into the number 42).
			*/
			result = any(total).(int) + any(number).(int)
		case float64:
			result = any(total).(float64) + any(number).(float64)
		case float32:
			result = any(total).(float32) + any(number).(float32)
		}
		total = result.(T)
	}
	return total
}

func main() {

	var v any
	v = Sum[string]("Jane", "Mark")
	fmt.Printf("%T\n", v)
	fmt.Println(v)
	v = Sum[int](1, 2, 3, 4, 5)
	fmt.Printf("%T\n", v)
	fmt.Println(v)
	v = Sum[float64](1.1, 2.2, 3.3, 4.4, 5.5)
	fmt.Printf("%T\n", v)
	fmt.Println(v)
	v = Sum[float32](1.1, 2.2, 3.3, 4.4, 5.5)
	fmt.Printf("%T\n", v)
	fmt.Println(v)
}
