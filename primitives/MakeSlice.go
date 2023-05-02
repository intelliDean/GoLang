package main

import "fmt"

func makeSlice() {
	slce := make([]int, 4, 9) //size and capacity
	slce[0] = 23
	slce[1] = 54
	slce[2] = 52
	slce[3] = 102
	fmt.Print(slce)
}

func main() {
	makeSlice()
}
