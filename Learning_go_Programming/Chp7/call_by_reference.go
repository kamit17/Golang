// address of variable is passed as an argument while calling a function.
package main

import "fmt"

func call_by_reference(x *int, y *int) {
	*x = *x - 20
	*y = *y - 30

}

func main() {
	var a int = 100
	var b int = 200
	fmt.Printf("Before calling fucntion, value of a : %d\n", a)
	fmt.Printf("Before calling fucntion, value of b : %d\n", b)

	call_by_reference(&a, &b)

	fmt.Printf("AFter calling fucntion, value of a : %d\n", a)
	fmt.Printf("after calling fucntion, value of b : %d\n", b)
}
