package main

import "fmt"

func main() {
	names := []string{"John", "Jane", "Jim", "Jill"}
	items := make([]int, 3, 10) // make is a built-in function that creates a slice of the given length and capacity
	fmt.Println(names)
	fmt.Println(items)
	fmt.Printf("Items: %+v, Length: %d, Capacity: %d\n", items, len(items), cap(items))
	items = append(items, 1)
	items = append(items, 2, 3, 4, 5)
	fmt.Printf("Items: %+v, Length: %d, Capacity: %d\n", items, len(items), cap(items))
	fmt.Println(items[3:])
}
