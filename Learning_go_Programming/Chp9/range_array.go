package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}
	sum := 0

	for _, v := range arr {
		sum = sum + v
	}
	fmt.Println("Sum of values", sum)

}
