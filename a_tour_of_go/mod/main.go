package main

import (
	"fmt"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"golang.org/x/tour/wc"
)

func main() {
	pic.Show(Pic)
	fmt.Print("\n\n")
	wc.Test(WordCount)

	reader.Validate(MyReader{})
}
