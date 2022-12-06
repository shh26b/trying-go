package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func main() {
	shihab := Person{"Shihab Mahamud", 19, "Faridpur"}

	fmt.Println(shihab)
	fmt.Println(shihab.address)
	fmt.Println(shihab.age)

	shihab2 := Person{name: "Shihab Again", age: 19}
	fmt.Println(shihab2)
	fmt.Println(shihab2.address)
	fmt.Println(shihab2.age)
}
