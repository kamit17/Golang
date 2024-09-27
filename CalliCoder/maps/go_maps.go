// A map is an unordered  collectin of key-value pairs.It maps keys to values. The keyrs are unique within a map while the values may not be.

// Declaring a map syntax
//var m map [KeyType]Value Type

//declaring a map of string keys to int values

//zero value of a map is nil . A nil map has no keys

package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}

}
