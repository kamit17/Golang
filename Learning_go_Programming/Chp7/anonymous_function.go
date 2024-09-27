/*A function without any name is called lambda or anonymous function.
These functions are assigned to a variable and that variable is used to call
the function.
*/

package main

import "fmt"

func return_anon() func(string) { //declare type of return value as  a function

	return func(message string) {
		fmt.Println(message)
	}

}

func main() {
	return_anon_variable := return_anon()

	fmt.Println(return_anon_variable) //prints the address of the function returned by 'return anon'
	return_anon_variable("Hello")
}
