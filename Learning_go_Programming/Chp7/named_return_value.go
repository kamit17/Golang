// Go function can return the name values as defined with the signature of a function.
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

	var c string
	result, c := sum(a, b, "Hello Go")
	fmt.Println("Values returned from the function are : ", result, c)
}
