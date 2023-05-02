package main

import "fmt"

func generateOdd() func() int {
	x := 1
	return func() (odd int) {
		odd = x
		x += 2
		return odd
	}
}
func main() {
	result := generateOdd()
	for i := 0; i < 5; i++ {
		fmt.Println(result())
	}

}
