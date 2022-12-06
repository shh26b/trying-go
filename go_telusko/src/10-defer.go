package main

import "fmt"

func main() {
	a()
}

func a() {
	fmt.Println("A begins")
	defer b()
	fmt.Println("A ends")
}

func b() {
	fmt.Println("B begins")
	fmt.Println("B ends")
}
