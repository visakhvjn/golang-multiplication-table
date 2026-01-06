package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Welcome to the Multiplcation Table Generator!")

	var number int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)

	if number > math.MaxInt64/10 || number < math.MinInt64/10 {
		fmt.Println("Please enter a valid number.")
		return
	}

	fmt.Printf("\nMultiplication table of %d is:\n", number)
	fmt.Println("--------------------------------")

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*int64(i))
	}

	fmt.Println("--------------------------------")
}
