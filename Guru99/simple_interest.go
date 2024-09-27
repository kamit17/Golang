package main

import (
	"fmt"
	"math"
)

func main() {

	var _principalAmount float64
	var _roi float64
	var _time int

	fmt.Print("Enter the principal amount: ")
	fmt.Scan(&_principalAmount)

	fmt.Print("Enter the rate of interest: ")
	fmt.Scan(&_roi)

	fmt.Print("Enter the duration: ")
	fmt.Scan(&_time)

	SI := (_principalAmount * _roi * float64(_time)) / 365

	fmt.Println("Simple interest is : ", math.Round(SI))
}
