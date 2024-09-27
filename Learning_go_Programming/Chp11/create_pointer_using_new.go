//new function is used to allocate memory space at runtime. It allocates memory accordig to the passed data type and returns its address toa pointer

package main

import "fmt"

func main() {
	a := new(int)
	fmt.Println(*a)
	*a = 10
	fmt.Println(*a)
}
