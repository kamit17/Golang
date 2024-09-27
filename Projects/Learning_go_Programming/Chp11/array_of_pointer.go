package main

import "fmt"

func main() {
	var a, b, c int
	var pointerv [4]*int
	pointerv[0] = &a
	pointerv[1] = &b
	pointerv[3] = &c
	fmt.Println(pointerv)
}
