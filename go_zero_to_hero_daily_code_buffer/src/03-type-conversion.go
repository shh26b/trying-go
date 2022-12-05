package main

import (
	"fmt"
	"strconv"
)

func main() {
	// type conversion
	i5 := 12
	var j float32 = float32(i5)
	j2 := float32(i5)

	fmt.Println(j, j2)

	// converting int to string
	var foo int = 97
	bar := string(rune(foo))
	buzz := strconv.Itoa(foo)

	fmt.Println(bar)
	fmt.Println(buzz)

	bar2 := true
	bar3 := false

	fmt.Printf("%v %T\n", bar2, bar2)
	fmt.Printf("%v %T\n", bar3, bar3)
}
