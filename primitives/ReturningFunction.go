package main

import "fmt"

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		i += 2
		ret = i
		return
	}
}
func main() {
	nextEven := makeEvenGenerator()
	for i := 0; i < 10; i++ {
		fmt.Print(nextEven(), " ")
	}
}
