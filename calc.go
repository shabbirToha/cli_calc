package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run calc.go [number1] [operator] [number2]")
		return
	}

	// Parse first number
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error: first argument is not a valid number")
		return
	}

	// Parse operator
	op := os.Args[2]

	// Parse second number
	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Error: third argument is not a valid number")
		return
	}

	var result float64

	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Error: unsupported operator. Use +, -, *, or /")
		return
	}

	fmt.Printf("%v %s %v = %v\n", num1, op, num2, result)
}
