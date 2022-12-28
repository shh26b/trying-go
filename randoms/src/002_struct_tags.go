package main

import (
	"fmt"
	"reflect"
	"strings"
)

type animal struct {
	Name string `required:"true" max-min:"10,3"`
	Age  int    ``
	F    string `r:"in good format"`
}

func prt(x interface{}) {
	fmt.Printf("Type = %T, Value = %+v\n", x, x)
}

func main() {
	a := animal{"cat", 10, "fy"}

	t := reflect.TypeOf(a)

	field, _ := t.FieldByName("Age")
	prt(field.Tag)

	name, _ := t.FieldByName("Name")
	prt(name.Tag)

	required := name.Tag.Get("required")
	maxMin := name.Tag.Get("max-min")
	prt(required)
	prt(maxMin)

	maxMinS := strings.Split(maxMin, ",")
	prt(maxMinS)

	max := maxMinS[0]
	min := maxMinS[1]
	prt(max)
	prt(min)

	f, _ := t.FieldByName("F")
	x, _ := f.Tag.Lookup("r")
	x2 := f.Tag.Get("r")
	prt(x)
	prt(x2)

	x3 := f.Tag.Get("f")
	x4, ok := f.Tag.Lookup("f")
	prt(x3)

	fmt.Println(ok)

	if ok {
		prt(x4)
	} else {
		fmt.Println("f field is empty")
	}
}
