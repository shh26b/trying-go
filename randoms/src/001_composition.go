package main

import "fmt"

type Comic struct {
	Universe string
}

func (c *Comic) ComicUniverse() string {
	return c.Universe
}

type Marvel struct {
	// anonymous field,
	// this is composition where
	// ths struct is embedded
	Comic
}

type DC struct {
	// anonymous field
	Comic
}

func main() {
	c1 := Marvel{Comic{Universe: "MCU"}}
	fmt.Println("Universe is: ", c1.ComicUniverse())

	c2 := DC{Comic{Universe: "DC"}}
	fmt.Println("Universe is: ", c2.ComicUniverse())

}
