/*
Variadic functions are flexible with passing arguments. These functions
can take an arbitrary number of arguments. It is necessary to specify the
type of the last parameter. The symbol specifies that the function can take
any number of arguments, whether zero or more
*/

package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args { // args is a list of passed arguments while calling the function.
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3))
}
