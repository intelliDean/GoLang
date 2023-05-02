package main

import "fmt"

func small(x []int) int {
	small := x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < small {
			small = x[i]
		}
	}
	return small
}

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	result := small(x)
	fmt.Print(result)
}
