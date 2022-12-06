package main

import "fmt"

func main() {
	var num1 int = 2
	var num2 = 3
	var result = num1 + num2
	fmt.Println(result)

	var defaultIntValue int
	fmt.Println(defaultIntValue) // 0

	var a, b int
	a, b = 10, 20

	fmt.Println(a, b)

	x := 100           // int
	y := 100.234       // float64
	c := 'a'           // rune or int32
	z := "Hello World" // string
	t := true || false // bool
	i := 3 + 5i        // complex128

	fmt.Println(x, y, c, z, t, i)
}
