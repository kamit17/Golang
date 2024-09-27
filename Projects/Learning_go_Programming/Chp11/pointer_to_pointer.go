package main

import "fmt"

func main() {
	var a int
	pointerv := &a
	*pointerv = 15
	fmt.Println("Value of pointerv: ", pointerv)
	fmt.Println("Address of pointerv:", &pointerv)

	pointerv1 := &pointerv
	fmt.Println("Value of pointerv1: ", pointerv1)
	fmt.Println("Address of pointerv1:", &pointerv1)

	pointerv2 := &pointerv1
	fmt.Println("Value of pointerv2: ", pointerv2)
	fmt.Println("Address of pointerv2:", &pointerv2)

	pointerv3 := &pointerv2
	fmt.Println("Value of pointerv3: ", pointerv3)
	fmt.Println("Address of pointerv3", &pointerv3)

}
