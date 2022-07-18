//The new() function takes a type as an argument, allocates enough memory to accommodate a value of that type, and returns a pointer to it.

package main

import "fmt"

func main() {

	ptr := new(int) //pointer to an 'int' type
	*ptr = 100

	fmt.Printf("Ptr = %#x,Ptr Value = %d\n", ptr, *ptr)
}
