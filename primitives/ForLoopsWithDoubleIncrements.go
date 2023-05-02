package main

import "fmt"

func forLoopWithDoubleIncrement() {
	for i, j := 0, 10; i < 10; i++ { // was supposed to use the i = i + 1, j = j - 1, but it's not working that way
		j--
		fmt.Println(i, " - ", j)
	}
}

func main() {
	forLoopWithDoubleIncrement()
}
