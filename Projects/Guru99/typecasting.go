package main

import "fmt"

func main() {
	var a int = 10
	var b int32 = 20
	var c byte = 15
	var d float32 = 0.05

	fmt.Println(int32(a) + int32(b))
	fmt.Println(int32(b) + int32(c))
	fmt.Println(int(a) + int(c))
	fmt.Println(float32(a) + float32(d))

}
