package main

import "fmt"

func useIf() {
	//guess until you guess right: fmt:Scanln(*number) worked instead of the format
	guess := 10
	var number int
	var cond bool = true
	for cond /*or true*/ {
		fmt.Print("Enter number to guess the lucky number: ")
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println("Invalid input, please try again")
			continue
		}
		if number > guess {
			fmt.Println("too high")
		} else if number < guess {
			fmt.Println("too low")
		} else {
			fmt.Println("Correct guess")
			cond = false
			/*or break*/
		}
	}
}

func main() {
	useIf()
}
