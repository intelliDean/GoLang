package main

import "fmt"

func averageArray() {
	numbers := [...]float64{594.0, 43.90, 2343, 5232, 434, 5554}
	total := 0.0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	fmt.Println(total / float64(len(numbers)))
}

func main() {
	averageArray()
}
