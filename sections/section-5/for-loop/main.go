package main

import "fmt"

func main() {
	// For loop is only way to iterate over a collection of data
	// C++ style for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("--------------------------------")

	// White style
	k := 4
	for k > 0 {
		fmt.Println(k)
		k--
	}

	fmt.Println("--------------------------------")

	j := 0
	for j >= 0 {
		j++
		if j%2 == 0 {
			continue
		}
		fmt.Println(j)

		if j > 10 {
			break
		}
	}

	fmt.Println("--------------------------------")

	items := [3]string{"item1", "item2", "item3"}
	for index, item := range items {
		fmt.Printf("Index: %d, Item: %s\n", index, item)
		fmt.Println(items[index])
	}
}
