package main

import "fmt"

type Student struct {
	name  string
	age   int
	phone string
	score int
}

func main() {
	s1 := Student{"Shihab", 19, "017", 95}
	fmt.Println(s1)

	s2 := Student{name: "Shihab", age: 19}
	fmt.Println(s2)
}
