//complex 64 : both real and imaginary parts of float 32 bytes
//complex 128: both real and imaginary parts are of float 64 type
//defatult type is complex128

package main

import "fmt"

func main() {
	var a = 3 + 5i
	var b = 2 + 4i
	var c = complex(a, b)

	var res1 = a + b
	var res2 = a - b
	var res3 = a * b
	var res4 = a / b

	fmt.Println(res1, res2, res3, res4)

}
