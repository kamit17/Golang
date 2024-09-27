package main

import "fmt"

func main() {
	var numbers [3]string //Declaring a string array of size 3 and adding elements
	numbers[0] = "One"
	numbers[1] = "Two"
	numbers[2] = "Three"
	fmt.Println(numbers[1])
	fmt.Println(len(numbers))
	fmt.Println(numbers)

	directions := [...]int{1, 2, 3, 4, 5}
	fmt.Println(directions)      //prints [1 2 3 4 5]
	fmt.Println(len(directions)) //prints 5

}
