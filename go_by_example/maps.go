package main

import (
	"fmt"
	"maps"
)

func main() {

	// create a map using builtin make(map[key-value]val-type)
	m := make(map[string]int)

	// set Key/Value pairs using name[key] = val syntac
	m["k1"] = 7
	m["k2"] = 13

	//Printing  a map will show all of its key value pairs
	fmt.Println("map: ", m)

	//Get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//if they key doesn't exist the zero value of the value type is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	//returning length of map
	fmt.Println("len:", len(m))

	//deleting key/value from map
	delete(m, "k2")
	fmt.Println("map:", m)

	//remove all key/value pairs from a map using clear builtin
	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//declare and initialize a map in same line

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map", n)

	//maps package contains a number of useful utility of functions for maps.

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
