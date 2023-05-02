package main

import "fmt"

func zeroWithPointer(input *int) {
	*input = 0
}

func zeroWithoutPointer(input int) {
	input = 0
}
func main() {
	x := 5
	y := 10
	zeroWithPointer(&x)
	zeroWithoutPointer(y)
	fmt.Println(x) // x is 0
	fmt.Println(y) // y is 10
}
