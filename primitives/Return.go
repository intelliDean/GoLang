package main

import "fmt"

func returnLiteral(x int, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	result := returnLiteral(23, 65)
	fmt.Print(result)
}
