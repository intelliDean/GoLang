package main

import "fmt"

func nestMap() {
	natives := map[string]map[string]string{
		"Cohort 13": map[string]string{
			"male":   "27",
			"female": "8",
		},
		"Cohort 14": map[string]string{
			"male":   "31",
			"female": "6",
		},
		"Cohort 15": map[string]string{
			"male":   "12",
			"female": "11",
		},
	}
	fmt.Println(natives["Cohort 13"])
	fmt.Print(natives["Cohort 13"]["male"])
}

func main() {
	nestMap()
}
