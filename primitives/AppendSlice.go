package main

import "fmt"

func appendSlice() {
	slces := make([]int, 4, 6)
	for i := 0; i < 10; i++ {
		slces[i] = (i + 1) * 23
		slces = append(slces, slces[i])
		if value := slces[i] == 46; value {
			fmt.Printf("%v : The slice contain %d\n", value, slces[i])
		}
	}
	fmt.Println(slces)
}

func main() {
	appendSlice()
}
