// pointers allow you to pass references to values and records within your program

package main

import "fmt"

//zeroval  has int parameter. So  arguments will be passed by value
//zeroval will get a copy of ival distinct from the one in the calling funciton
func zeroval(ival int) {
	ival = 0
}

//zeroptr has *int parameter(int pointer).
//*iptr dereferecnes the pointer from its memory address to the current value at the address.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i) //&i syntax gives memory address of i
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
