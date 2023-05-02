package main

import "fmt"

func forLoopWithArray() {
	a := [4]int{21, 42, 23, 49} //array
	for k, v := range a {
		fmt.Println(k, v)
	}
}

func main() {
	forLoopWithArray()
}
