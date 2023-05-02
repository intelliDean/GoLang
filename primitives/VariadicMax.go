package main

/*Write a function with one variadic parameter
that finds the greatest number in a list of numbers.*/

func variadic(numbers ...int) int {
	max := numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max
}

func main() {
	println(variadic(23, 54, 12, 54, 34, 89, 65, 67))
}
