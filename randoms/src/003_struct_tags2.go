package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	F1 string `json:"field1"`
	F2 string `json:"field2,omitempty"`
	F3 string `json:"field3,omitempty"`
	F4 string `json:"-"`
}

func prt(x interface{}) {
	fmt.Printf("Type = %T, Value = %+v\n", x, x)
}

func main() {
	t := T{"v1", "", "v2", "v3"}

	fmt.Println(t)

	sb, _ := json.Marshal(t)
	prt(sb)

	s := string(sb)
	prt(s)
}
