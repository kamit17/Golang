// Dereferencing  is basically getting the value stored in an address pointed by the pointer

package main

import "fmt"

func main() {
	a := 10
	var pointerv *int
	pointerv = &a
	fmt.Println("Address of a is ", pointerv)
	fmt.Println("Value of a is ", *pointerv)
}
