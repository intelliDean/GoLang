package main

/*The Fibonacci sequence is defined as: fib(0) =
0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
Write a recursive function which can find
fib(n)*/

func fibo(number int) int {
	if number == 1 {
		return 1
	}
	if number == 0 {
		return 0
	}
	return fibo(number-1) + fibo(number-2)
}

func main() {
	println(fibo(10))
}
