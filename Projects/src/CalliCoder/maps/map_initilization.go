//initalizing a map using the biuld-in make() function
// var m = make(map[string]int)

package main

import "fmt"

func main() {
	var m = make(map[string]int)

	fmt.Println(m)

	if m == nil {
		fmt.Println("m is nil")

	} else {
		fmt.Println("m is not nil")
	}

	//make() function returns an initialized and ready to use map.
	//Since it is initialized, you can add new keys to it

	m["one hunderd"] = 100
	fmt.Println(m)

	// initializzing a  mapusing a map literal
	// pass the key - value pairs seperated by colong inside curly braces
	var map_literal = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	fmt.Println(map_literal)

}
