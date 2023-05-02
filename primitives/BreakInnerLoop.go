package main

import "fmt"

func breakLoop() {
Loop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i * j)
			if i*j > 3 {
				break Loop
			}
		}
	}
}

func main() {
	breakLoop()
}
