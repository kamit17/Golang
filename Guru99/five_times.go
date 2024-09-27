package main

import "fmt"

func main() {
	var int1, int2, int3, int4, int5 int
	fmt.Print("Enter the first integer: ")
	fmt.Scan(&int1)
	fmt.Print("Enter the second integer: ")
	fmt.Scan(&int2)
	fmt.Print("Enter the third integer: ")
	fmt.Scan(&int3)
	fmt.Print("Enter the fourth integer: ")
	fmt.Scan(&int4)
	fmt.Print("Enter the fifth integer: ")
	fmt.Scan(&int5)

	sum := int1 + int2 + int3 + int4 + int5
	fmt.Println("Sum of the integers is: ", sum)

	product := int1 * int2 * int3 * int4 * int5
	fmt.Println("Product: ", product)

	average := float64(sum) / 5
	fmt.Println("average: ", average)

}
