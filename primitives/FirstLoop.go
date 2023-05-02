package main

import "fmt"

func firstLoop() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("%v ", i)
		if i%20 == 0 {
			fmt.Println()
		}
	}
}

func main() {
	firstLoop()
}
