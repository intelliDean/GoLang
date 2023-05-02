package main

import (
	"fmt"
	"math/rand"
)

func testSlice() {
	var numbers [11]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(12)
		fmt.Printf("Number %d: %d\n", i+1, numbers[i])
	}
	fmt.Println("there are ", len(numbers), " numbers in the slice")
}

func main() {
	testSlice()
}
