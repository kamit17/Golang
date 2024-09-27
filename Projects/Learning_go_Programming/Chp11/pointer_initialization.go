//initializing pointers means assigning an address of a variable to the pointer which is denoted by &

package main

import "fmt"

func main() {
	//var p *int
	i := 42
	var pointerv *int

	pointerv = &i
	fmt.Println(pointerv)
	fmt.Println(*pointerv)
	fmt.Println(&pointerv)
}
