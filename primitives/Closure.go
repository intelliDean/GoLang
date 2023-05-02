package main

import "fmt"

func nestedMethod(num1 int, num2 int) int {
	sumResult := func(num1 int, num2 int) int {
		return num1 + num2
	}
	return sumResult(num1, num2)
}

func main() {
	fmt.Print("Enter first number: ")
	var num1 int
	_, err := fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	var num2 int
	_, err2 := fmt.Scan(&num2)

	if err != nil && err2 != nil {
		fmt.Print("Illegal input")
		return
	}

	result := nestedMethod(num1, num2)

	fmt.Print("Enter another first number: ")
	var num3 int
	_, err3 := fmt.Scan(&num1)
	fmt.Print("Enter another second number: ")
	var num4 int
	_, err4 := fmt.Scan(&num2)

	if err3 != nil && err4 != nil {
		fmt.Print("Illegal input")
		return
	}
	anotherSum := func(number1 int, number2 int) int {
		return number1 + number2
	}

	newResult := anotherSum(num3, num4)

	fmt.Print(newResult + result) //this is not working as it should yet, I will fix it later
}
