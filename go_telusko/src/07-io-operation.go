package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	defer writer.Flush()

	n := 0
	scanf("%d", &n)

	var arr [100]int

	for i := 0; i < n; i++ {
		scanf("%d", &arr[i])
	}

	for i := 0; i < n; i++ {
		printf("%d ", arr[i])
	}
}
