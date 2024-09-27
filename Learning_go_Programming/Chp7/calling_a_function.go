package main

import "fmt"

func sum(a, b int, c string) (int, string) {
	sum_value := a + b
	return sum_value, c
}

func main() {
	a := 30
	b := 90
	var c string
	result, c := sum(a, b, "Hello Go")
	fmt.Println("values returned from the function are : ", result, c)
}
