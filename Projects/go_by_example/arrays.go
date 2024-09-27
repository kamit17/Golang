package main

import "fmt"

func main() {

	var a [5]int // creating an array to hold exacly 5 ints  . By deafult an array is zero valued.
	fmt.Println("emp:", a)

	a[4] = 100 // set value using array[index]
	fmt.Println("set:", a)
	fmt.Println("get:", a[4]) // get value with array[index]

	fmt.Println("len:", len(a)) // built in len returns lenght of array

	b := [5]int{1, 2, 3, 4, 5} // declare and initialize an array in one line
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

}
