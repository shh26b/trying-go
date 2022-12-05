package main

import "fmt"

// global variables
var i = 32
var I = "exporting variable"

func main() {
	fmt.Printf("%v <- %T\n", i, i)

	// declaring variable
	var i int
	i = 12

	var i2 int = 14

	var i3 = 16

	i4 := 18

	fmt.Printf("%v <- %T\n", i, i)
	fmt.Printf("%v <- %T\n", i2, i2)
	fmt.Printf("%v <- %T\n", i3, i3)
	fmt.Printf("%v <- %T\n", i4, i4)

	// default values
	var x int
	var x2 float32
	var x3 string

	fmt.Printf("%#v %#v %#v\n", x, x2, x3)

	// default type inference
	integer_value := 5
	float_value := 23.55
	string_value := "Hello world"

	fmt.Printf("%v <- %T\n", integer_value, integer_value)
	fmt.Printf("%v <- %T\n", float_value, float_value)
	fmt.Printf("%v <- %T\n", string_value, string_value)
}
