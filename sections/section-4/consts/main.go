package main

import "fmt"

const GREETING = "Hello World!"
const COUNT = 10

func main() {
	fmt.Println(GREETING)
	fmt.Println(COUNT)

	const pi float64 = 3.14
	fmt.Println(pi)
}
