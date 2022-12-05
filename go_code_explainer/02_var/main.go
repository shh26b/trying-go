package main

import "fmt"

const pi = 3.1416

func main() {
	// Defining vars
	var i int
	var j int = 1
	var k = 2
	var l, m = 3, 4
	n := 5
	fmt.Println(i, j, k, l, m, n)

	// Literals and default type in go
	a := 5     // int - Integer literals
	b := 5.5   // float64 - Floating-point literals
	c := 1.23i // complex128 -  Imaginary literals
	d := 'c'   // rune - Rune literals
	e := "str" // string - String literals
	f := true  // bool - Boolean literals
	fmt.Println(a, b, c, d, e, f)

	// default value
	var a2 int        // 0
	var b2 float64    // 0.0
	var c2 complex128 // 0+0i
	var d2 rune       // 0
	var e2 string     // ""
	var f2 bool       // false
	fmt.Println(a2, b2, c2, d2, e2, f2)

	fmt.Println(pi)
}
