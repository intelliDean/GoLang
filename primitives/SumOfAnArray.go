package main

import "fmt"

func sumOfArrayElements() {
	numbers := [...]int{23, 89, 54, 120, 432}
	var total int
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	//for _, number := range numbers {
	//	total += number
	//}
	fmt.Println(total)
}

func main() {
	sumOfArrayElements()
}
