package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstNumber string
	var secondNumber string

	var firstInt int
	var secondInt int

	//fmt.Println("Enter the first integer: ")
	//fmt.Println("Enter the second integer: ")

	fmt.Print("Enter the first integer: ")
	fmt.Scan(&firstNumber)

	fmt.Print("Enter the second integer: ")
	fmt.Scan(&secondNumber)

	firstInt, _ = strconv.Atoi(firstNumber)
	secondInt, _ = strconv.Atoi(secondNumber)

	//fmt.Println(firstNumber + secondNumber)

	fmt.Println(firstInt + secondInt)

}
