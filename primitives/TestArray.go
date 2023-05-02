package main

import "fmt"

func testArray() {
	var ages [8]int
	var eachAge int
	for i := 0; i < len(ages); {
		fmt.Print("Enter age at index ", i+1, ": ")
		_, err := fmt.Scan(&eachAge)
		if err != nil {
			fmt.Print("Invalid input: Enter integer\n")
			continue
		}
		ages[i] = eachAge
		i++
	}
	fmt.Println(ages)
}

func main() {
	testArray()
}
