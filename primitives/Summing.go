package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func main() {
	number1 := 23
	number2 := 43
	result := sum(number1, number2)
	fmt.Print(result)
}
