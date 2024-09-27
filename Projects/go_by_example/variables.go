/*
Variables are explicitly declared and used by the compiler

*/

package main

import "fmt"

func main() {
	var a = "initial" // Var declares 1 or more variables
	fmt.Println(a)

	var b, c int = 1, 2 // you can declare multiple variables at once.
	fmt.Println(b, c)

	var d = true //Go will infer the type of intialzed variable
	fmt.Println(d)

	var e int
	fmt.Println(e) // Variabled declared without a corresponding initialization are zero valued

	f := "apple" //:= is shorthand for declaring and initializing a variable
	fmt.Println(f)
}
