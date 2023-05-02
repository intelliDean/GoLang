package main

func multiple(x []int) int {
	total := 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	return total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	println(multiple(numbers))

}
