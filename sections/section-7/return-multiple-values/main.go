package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func getStats(a, b int) (int, float64, string, error) {
	if b == 0 {
		return 0, 0, "", fmt.Errorf("division by zero")
	}
	quotient := a / b
	average := float64(a+b) / 2.0
	msg := "ok"
	return quotient, average, msg, nil
}

func main() {
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", res)
	}

	res, err = divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", res)
	}

	quotient, average, msg, err := getStats(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient)
		fmt.Println("Average:", average)
		fmt.Println("Message:", msg)
	}
}
