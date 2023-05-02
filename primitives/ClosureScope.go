package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	for i := 0; i < 10; i++ {
		fmt.Println(increment())
	}
}
