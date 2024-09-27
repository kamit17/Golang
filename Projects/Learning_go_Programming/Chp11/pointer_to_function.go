package main

import "fmt"

func main() {

	a := "Welcome to the Jungle"
	pointerv := &a
	fmt.Println("value if a is : ", pointerv)
	changeVal(a)
	fmt.Println("After modification, value of a is:", *pointerv)
	changeVal1(pointerv)
	fmt.Println("After modification, value of a is:", *pointerv)

}

func changeVal(y string) {
	y = "We got one endgame"
	fmt.Println("Changed value is :", y)
}

func changeVal1(y *string) {
	*y = "We got one endgame"
	fmt.Println("Changed value is :", y)
}
