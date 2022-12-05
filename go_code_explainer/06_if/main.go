package main

import "fmt"

func main() {
	x := 5

	if x%2 == 0 {
		fmt.Println(x, "is even")
	} else if x <= 50 {
		fmt.Println(x, "is odd and less than and equal 50")
	} else {
		fmt.Println(x, "is odd")
	}

	bird := "king"

	switch bird {
	case "parrot":
		fmt.Println("parrot have green feathers")
	case "eagle":
		fmt.Println("Eagle have brown feathers")
	default:
		fmt.Println("Default")
	}
}
