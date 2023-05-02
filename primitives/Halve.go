package main

import "fmt"

/*
	Write a function which takes an integer and

halves it and returns true if it was even or false
if it was odd. For example half(1) should return
(0, false) and half(2) should return (1,
true).
*/
func halve(number int) (int, bool) {
	value := number / 2
	if out := value%2 == 0; out {
		return value, out
	} else {
		return value, out
	}
}

func main() {
	fmt.Print(halve(8))
}
