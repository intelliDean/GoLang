package main

import "fmt"

func specialFunction(input ...int) (total int) {
	total = 0
	for i := 0; i < len(input); i++ {
		total += input[i]
	}
	return
}

func main() {
	fmt.Printf("%d", specialFunction(32, 54, 28, 437, 236, 83, 274)) // this takes zero or more parameters
}
