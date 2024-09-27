package main

import "fmt"

func main() {
	var arrayval [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := arrayval[0:5]
	slice2 := slice1[:3]

	fmt.Println("Slice from array:", slice1)
	fmt.Println("Slice from slice:", slice2)

}
