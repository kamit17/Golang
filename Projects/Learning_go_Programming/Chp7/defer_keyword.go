// Defer keyword is used to execute a functino at the end, after executing all the statements specified in the main function.

package main

import "fmt"

func defer_fun() {
	fmt.Println("Defered function")
}
func main() {
	fmt.Println("Before defer statement")
	defer defer_fun()
	fmt.Println("After defer statement")
}
