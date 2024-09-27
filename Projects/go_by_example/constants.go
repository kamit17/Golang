// Go supports char, string, boolean and numberic values

package main

import (
	"fmt"
	"math"
)

const s string = "constant" // const declared a constant

func main() {
	fmt.Println(s)

	const n = 50000000 // A const statement can appear anywhere a var statement can

	const d = 3e20 / n // Constant expressions perform arithmetic with arbitrary prevision
	fmt.Println(d)

	fmt.Println(int64(d)) // A numeric constant has not type until it is given one

	fmt.Println(math.Sin(n)) // A number can be given a type by using it in a context that requires one, such as a variable assignment or function
}
