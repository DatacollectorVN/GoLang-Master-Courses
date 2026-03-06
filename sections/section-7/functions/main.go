package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a, b int) int {
	return a + b
}

// recursive function
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// closure function
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	greet("Nathan")
	res := add(1, 2)
	fmt.Println("Result:", res)

	res = factorial(5)
	fmt.Println("Factorial:", res)

	nextInt := intSeq()
	fmt.Println("Next int:", nextInt())
	fmt.Println("Next int:", nextInt())

	sum := adder()
	fmt.Println("Sum:", sum(1))
	fmt.Println("Sum:", sum(2))
	fmt.Println("Sum:", sum(3))
	fmt.Println("Sum:", sum(4))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	logger := func(message string) {
		fmt.Println("LOG:", message)
	}
	logger("This is a log message")
	logger("This is a log message")
	logger("This is a log message")
}
