package main

import "fmt"

func name() (string, int) {
	var check string = "my name is Michael Dean"
	var age int = 23
	return check, age
}

func main() {
	intro, age := name()

	fmt.Println(intro)

	fmt.Printf("and I am %d", age)
}
