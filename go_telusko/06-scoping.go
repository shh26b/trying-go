package main

import "fmt"

var b = "package scope variable"

func demo() {
	a := "function scope variable"
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	demo()
	fmt.Println(b)
}
