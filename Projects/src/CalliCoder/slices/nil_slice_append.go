package main

import "fmt"

func main() {
	var s []string

	//Appending to a nil slice
	s = append(s, "cat", "dog", "Lion", "Tiger")

	fmt.Printf("s = %v,len = %d,cap = %d\n", s, len(s), cap(s))
}
