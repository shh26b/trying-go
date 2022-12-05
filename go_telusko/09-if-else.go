package main

import "fmt"

func main() {
	num := 8

	if num == 0 {
		fmt.Println(num, "is zero")
	} else if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	switch num {
	case 0:
		fmt.Println(num, "is zero")
	default:
		fmt.Println(num, "is non zero")
	}
}
