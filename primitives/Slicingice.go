package main

import "fmt"

func slicingSlice() {
	numbers := []int{23, 54, 23, 89, 43, 109, 531, 90}
	fmt.Println(numbers[0:5])
	fmt.Println(numbers[0:])
	fmt.Println(numbers[:])
	fmt.Println(numbers[3:len(numbers)])
	fmt.Println(numbers[:len(numbers)])
}

func main() {
	slicingSlice()
}
