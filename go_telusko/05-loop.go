package main

import "fmt"

func main() {
	// for {
	// 	fmt.Println("Hello")
	// }

	i := 1
	for i <= 5 {
		fmt.Println("Hello", i)
		i += 1
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("Hello", i)
	}
}
