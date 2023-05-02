package main

import "fmt"

func subArray() int {
	array := [7]int{4, 9, 12, 13, 2, 1, 0}
	max := 0
	inc := 3
	for i := 0; i < len(array)-3; i++ {
		k := i
		count := 0
		total := 0
		for j := 0; j < inc; j++ {
			total += array[k]
			k++
			count++
			if total > max {
				max = total
			}
		}
		inc++
	}
	return max
}

func main() { //throwing errors, I will fix it later
	result := subArray()
	fmt.Print(result)
}
