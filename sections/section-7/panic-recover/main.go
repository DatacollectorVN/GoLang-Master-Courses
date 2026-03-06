package main

import "fmt"

/*
Panic is used to stop the execution of a function and start the recovery process.
It is often used to handle errors that cannot be recovered from.
It is also used to stop the execution of a function and start the recovery process.
*/

func mightPanic(is_panic bool) {
	if is_panic {
		panic("This is a panic")
	} else {
		fmt.Println("This is a function that might panic after panic is called")
	}
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	mightPanic(true)
}

func main() {
	recoverable()
}
