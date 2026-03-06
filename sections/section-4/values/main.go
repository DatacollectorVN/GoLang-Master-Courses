package main

import "fmt"

func main() {
	fmt.Printf("Hello World!")
	fmt.Println("Hello World!")
	fmt.Println("Hello" + " " + "World!")

	fmt.Printf("%+v\n", 100)
	fmt.Printf("%+v\n", []int{1, 2, 3})
	fmt.Printf("%+v\n", []string{"Hello", "World"})
}
