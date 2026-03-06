package main

import "fmt"

func sum(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func config(numbers ...int) int {
	if len(numbers) == 0 {
		fmt.Println("No numbers provided")
		return 0
	}
	return numbers[0]
}

func main() {
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)

	config := config(1, 2, 3, 4, 5)
	fmt.Println("Config:", config)
}
