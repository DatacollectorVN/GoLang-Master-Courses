package main

import "fmt"

func main() {
	var numbers [5]int // Zero value is [0 0 0 0 0]

	fmt.Println(numbers)

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	fmt.Println(numbers)

	numbers2 := [3]int{1, 2, 3}
	fmt.Println(numbers2)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	fmt.Println("--------------------------------")

	sum := 0
	for _, value := range numbers {
		sum += value
	}

	fmt.Println("Sum:", sum)

	var matrix [3][3]int // Zero value is [[0 0 0] [0 0 0] [0 0 0]]
	fmt.Println(matrix)

	fmt.Println("--------------------------------")

	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	matrix[2][0] = 7
	matrix[2][1] = 8
	matrix[2][2] = 9
	fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
