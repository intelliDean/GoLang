package main

import "fmt"

func slicing() {
	arrays := []int{12, 43, 76, 41, 54}
	arrays = append(arrays, 400, 87, 87, 32)
	fmt.Println(arrays)
}

func main() {
	slicing()
}
