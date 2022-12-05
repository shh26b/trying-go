package main

import "fmt"

func add_sub(x, y int) (int, int) {
	res := x + y
	res2 := x - y
	return res, res2
}

func add_sub2(x, y int) (res, res2 int) {
	res = x + y
	res2 = x - y
	return res, res2
}

func main() {
	num1 := 5
	num2 := 4

	result1, result2 := add_sub(num1, num2)
	result3, result4 := add_sub2(num1, num2)

	fmt.Println(result1, result2)
	fmt.Println(result3, result4)
}
