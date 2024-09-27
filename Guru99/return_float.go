package main

import (
	"fmt"
	"math"
)

func main() {
	var _passed_number float64
	fmt.Print("Enter a float number number: ")
	fmt.Scan(&_passed_number)

	fmt.Println("Integer part of the number passed is: ", math.Trunc(_passed_number))
}
