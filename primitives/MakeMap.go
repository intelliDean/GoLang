package main

import "fmt"

func ageMap() {
	ages := make(map[string]int)
	ages["Dean"] = 23
	ages["Mike"] = 12
	ages["John"] = 90
	ages["King"] = 43
	ages["Seun"] = 45

	if value, isPart := ages["John"]; isPart {
		fmt.Println(value, isPart)
	} else {
		fmt.Print("Unavailable")
	}
}

func main() {
	ageMap()
}
