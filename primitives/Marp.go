package main

import "fmt"

func marping() {
	marp := make(map[string]int)
	marp["First"] = 23
	marp["Second"] = 43
	marp["Third"] = 89
	marp["Fourth"] = 12
	marp["Fifth"] = 97
	for k, v := range marp {
		fmt.Println(k, v)
	}
	key, ok := marp["First"]
	fmt.Print(key, ok)
}

func main() {
	marping()
}
