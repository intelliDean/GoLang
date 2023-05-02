package main

import "fmt"

func forLoopWithSlice() {
	s := []int{1, 2, 3, 4} //slice
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}
}

func main() {
	forLoopWithSlice()
}
