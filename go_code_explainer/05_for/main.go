package main

import "fmt"

func facto(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	x := facto(10)
	fmt.Println(x)

	ans := 0
	i := 1
	for i <= 5 {
		ans += i
		i++
	}
	fmt.Println(ans)

	fact := []int{1, 2, 6, 24, 120}
	ln := len(fact)
	for i := 0; i < ln; i++ {
		fmt.Println(fact[i])
	}

	for i, val := range fact {
		fmt.Println(i, val)
	}

	for _, val := range fact {
		fmt.Println(val)
	}
}
