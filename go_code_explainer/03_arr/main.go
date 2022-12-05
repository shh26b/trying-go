package main

import "fmt"

func main() {
	// array
	var arr [5]int
	arr[1] = 34
	fmt.Println(arr)

	var arr2d [3][5]int
	arr2d[2][1] = 3
	fmt.Println(arr2d)

	fmt.Println(len(arr), len(arr2d))

	fact1 := [5]int{1, 2, 3, 6, 24}
	fmt.Println(fact1)

	// slice
	cut1 := fact1[1:3]
	cut2 := fact1[:3]
	cut3 := fact1[1:]
	cut4 := fact1[:]
	fmt.Println(cut1, cut2, cut3, cut4)

	cut1[0] = 100
	fmt.Println("fact1: ", fact1)
	fmt.Println("cut1: ", cut1)

	var sl []int
	fmt.Println(sl, len(sl))

	sl2 := make([]int, 5)
	sl2[2] = 35
	sl2 = append(sl2, 50, 23)
	fmt.Println(sl2, len(sl2))
}
