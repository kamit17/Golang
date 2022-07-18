// A pointer van be initialized with the memory address of another variable. The address of the variable can be retreived usng the & operator

package main

import "fmt"

func main() {
	var a = 5.67
	var p = &a

	fmt.Println("Value stored in variable a = ", a)
	fmt.Println("Address of variable a = ", &a)
	fmt.Println("Value stored in variable p = ", p)
	fmt.Println("Address of variable p = ", &p)
	fmt.Println("*p = ", *p)
}
