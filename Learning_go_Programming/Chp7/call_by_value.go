// the actual values of the variables is passed to a function.

package main

import "fmt"

func call_by_value(x, y int) {
	x = x - 20
	y = y - 30
	fmt.Printf("In Function, value of a : %d\n", x) //prints 80
	fmt.Printf("In function, value of  b : %d\n", y)

}

func main() {
	var a int = 100
	var b int = 200
	fmt.Printf("Before calling fucntion, value of a : %d\n", a)
	fmt.Printf("Before calling fucntion, value of b : %d\n", b)

	call_by_value(a, b)
	fmt.Printf("AFter calling fucntion, value of a : %d\n", a)
	fmt.Printf("after calling fucntion, value of b : %d\n", b)
}
