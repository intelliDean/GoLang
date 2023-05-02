package main

import "fmt"

func useSwitch() {
	var number int
	for true {
		fmt.Print("Enter number: ")
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}
		switch number {
		case 0:
			fmt.Print("Zero\n")
		case 1:
			fmt.Print("One\n")
		case 2:
			fmt.Print("Two\n")
		case 3:
			fmt.Print("Three\n")
		case 4:
			fmt.Print("Four\n")
		case 5:
			fmt.Print("Five\n")
		case 6:
			fmt.Print("Six\n")
		case 7:
			fmt.Print("Seven\n")
		case 8:
			fmt.Print("Eight\n")
		case 9:
			fmt.Print("Nine\n")
		default:
			fmt.Print("Above basic\n")
			return
		}
	}
}

func main() {
	useSwitch()
}
