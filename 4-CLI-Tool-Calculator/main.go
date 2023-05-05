package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator [number] [operator] [number]")
		os.Exit(1)
	}

	operand1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid first operand:", os.Args[1])
		os.Exit(1)
	}

	operator := os.Args[2]

	operand2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid second operand:", os.Args[3])
		os.Exit(1)
	}

	var result float64

	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "x":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Cannot divide by zero")
			os.Exit(1)
		}
		result = operand1 / operand2
	default:
		fmt.Println("Invalid operator:", operator)
		os.Exit(1)
	}

	fmt.Println("Result:", result)
}
