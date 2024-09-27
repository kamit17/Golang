package main

import "fmt"

func main() {
	var x [5]int // An array of 5 integers

	x[0] = 100
	x[1] = 101
	x[3] = 103
	x[4] = 105
	fmt.Printf("x[0] = %d, x[1] = %d, x[2] = %d, x[3]= %d\n", x[0], x[1], x[2], x[3])
	fmt.Println("x = ", x)

	var y [8]string //An array of 8 strings

	fmt.Println(y)

	var z [3]complex128 //An array of 3 compels numbers
	fmt.Println(z)

	// Declaring and initializing an array at the same time
	var a = [5]int{2, 4, 6, 8, 10}
	fmt.Println(a)

	// Short declaration for declaring and initializing an array
	b := [5]int{2, 4, 6, 8, 10}
	fmt.Println(b)

	// You don't need to initialize all the elements of the array.
	// The un-initialized elements will be assigned the zero value of the corresponding array type
	c := [5]int{2}
	fmt.Println(c)

	// Letting Go compiler infer the length of the array
	j := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println(j)

	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`

	a2[1] = "German"

	fmt.Println("a1 = ", a1) // The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)
}
