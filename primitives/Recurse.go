package main

import "fmt"

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	// Input a number to find its factorial
	var n uint
	fmt.Print("Enter a non-negative integer to calculate its factorial: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid number")
		return
	}

	// Calculate and print the factorial of n
	if n < 0 {
		fmt.Println("Error: Entered number is negative. Factorial can be calculated only for non-negative integers.")
	} else {
		fmt.Printf("The factorial of %d is %d\n", n, factorial(n))
	}
}
