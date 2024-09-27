//blank identifier are used when  we don't want some values returned by the function to be used further in our program.
// Blank identifier tells the compiler that this particular variable is not going to be used in the program.

// calling function sum, but we don't need the variable c, so we are using blank identifier for it.

package main

import "fmt"

func sum(a, b int, c string) (sum_value int, c_val string) {
	sum_value = a + b
	c_val = c
	return
}

func main() {
	a := 90
	b := 30
	result, _ := sum(a, b, "Hello Go")
	fmt.Println("Values returned from the function are : ", result)
}
