package main

import "fmt"

func main() {
	a := 10
	var pointerv *int
	fmt.Println("Value of pointer is ", pointerv)
	pointerv = &a
	fmt.Println("Location of a is ", &pointerv)
	fmt.Println("Value of a is ", *pointerv)
	*pointerv = 30
	fmt.Printf("After assigning new value,")
	fmt.Println("value of a is ", *pointerv)
}
