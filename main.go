package main

import (
	"fmt"
	"math"
)

// IsValidNumber checks if the number is within safe multiplication range
func IsValidNumber(number int64) bool {
	return number <= math.MaxInt64/10 && number >= math.MinInt64/10
}

// GenerateMultiplicationTable returns the multiplication table for a given number
func GenerateMultiplicationTable(number int64) []int64 {
	result := make([]int64, 10)
	for i := 1; i <= 10; i++ {
		result[i-1] = number * int64(i)
	}
	return result
}

// Multiply returns the product of two numbers
func Multiply(a, b int64) int64 {
	return a * b
}

func main() {
	fmt.Println("Welcome to the Multiplcation Table Generator!")

	var number int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if !IsValidNumber(number) {
		fmt.Println("Please enter a valid number.")
		return
	}

	fmt.Printf("\nMultiplication table of %d is:\n", number)
	fmt.Println("--------------------------------")

	table := GenerateMultiplicationTable(number)
	for i, result := range table {
		fmt.Printf("%d x %d = %d\n", number, i+1, result)
	}

	fmt.Println("--------------------------------")
}
