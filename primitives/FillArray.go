package main

import "fmt"

func fillArray() {
	arrays := [...]int{23, 34, 54, 89, 90}
	for i, array := range arrays {
		fmt.Println(i, " -> ", array)
	}
}

func main() {
	fillArray()
}
