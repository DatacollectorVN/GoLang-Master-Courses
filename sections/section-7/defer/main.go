package main

/*
Defer is used to delay the execution of a function until the surrounding function returns.
It is often used to ensure that a resource is released after the function completes.
It is also used to ensure that a function is called in a specific order.
*/
import (
	"fmt"
	"os"
)

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: deferred")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")
	fmt.Println("Function simpleDefer: Middle")

}

func lifoSimpleDefer() {
	fmt.Println("Function lifoSimpleDefer: Start")
	defer fmt.Println("First: deferred")
	defer fmt.Println("Second: deferred")
	fmt.Println("Function lifoSimpleDefer: Middle")
}

func main() {
	defer fmt.Println("deferred function")
	fmt.Println("main function")

	arr := []int{1, 2, 3, 4, 5}
	for i, v := range arr {
		if i == 2 {
			defer fmt.Println(v)
		}
		fmt.Println(v)
	}
	defer fmt.Println("deferred function 2")

	file, err := os.Create("./defer.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//simpleDefer()
	lifoSimpleDefer()

	fmt.Println("Last in main()")
}
