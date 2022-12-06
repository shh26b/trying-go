package main

import "fmt"

var memo [55]int

func fib(n int) int {
	if n <= 1 {
		return n
	}
	if memo[n] == -1 {
		memo[n] = fib(n-1) + fib(n-2)
	}
	return memo[n]
}

func main() {
	for i := 0; i < 55; i++ {
		memo[i] = -1
	}
	res := fib(50)

	fmt.Println(res)

	pre1, pre2 := 0, 1
	for i := 2; i <= 10; i++ {
		pre1, pre2 = pre2, pre1+pre2
	}
	fmt.Println(pre2)
}
