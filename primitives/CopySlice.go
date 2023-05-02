package main

import "fmt"

func copySlice() {
	slce := []int{43, 65, 23, 87}
	anotherSlce := make([]int, 3)
	copy(anotherSlce, slce) //keyword copy copies one object into another object of the same type
	fmt.Print(anotherSlce)
}

func main() {
	copySlice()
}
