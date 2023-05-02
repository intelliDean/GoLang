package main

import "fmt"

func returnMultiple(name string, age int) (returnedName string, returnedAge int) {
	returnedAge = age
	returnedName = name
	return
}

func main() {
	name, age := returnMultiple("Michael Dean", 23)
	fmt.Println(name)
	fmt.Print(age)
}
