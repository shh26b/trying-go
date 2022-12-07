package main

import (
	"fmt"

	a_our_exercises "github.com/shh26b/go-practice/mod-project/a_tour_exercises"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
)

func main() {
	pic.Show(a_our_exercises.Pic)
	fmt.Print("\n\n")
	wc.Test(a_our_exercises.WordCount)
}
