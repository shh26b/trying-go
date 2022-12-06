package main

import "fmt"

func main() {
	mp := make(map[string]float64)
	mp["go"] = 1.19
	mp["py"] = 3.11
	mp["java"] = 1.19
	mp["node"] = 16.17
	fmt.Println(mp)

	mp2 := map[string]float64{"go": 1.19, "py": 3.11}
	fmt.Println(mp2)

	fmt.Println(mp["go"], mp["node"], mp["c++"])
	val, ok := mp["code"]
	if ok {
		fmt.Println(val, "is present")
	} else {
		fmt.Println(val, "is not present")
	}

	val2, ok2 := mp["go"]
	if ok2 {
		fmt.Println(val2, "is present")
	} else {
		fmt.Println(val2, "is not present")
	}

	mp["go"] = 1.20
	fmt.Println(mp)

	delete(mp, "go")
	fmt.Println(mp)
}
