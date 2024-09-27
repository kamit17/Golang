package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64 = -30.4, 45.6, 4
	fmt.Println("a =", a, "\nb=", b, "\nc=", c)

	fmt.Println("math.Abs(a) =", math.Abs(a))
	fmt.Println("math.Ceil(b) = ", math.Ceil(b))
	fmt.Println("math.Floor(a) = ", math.Floor(a))
	fmt.Println("math.Exp(a) = ", math.Exp(a))
	fmt.Println("math.Sqrt(c) = ", math.Sqrt(c))
	fmt.Println("math.Trunc(a) = ", math.Trunc(a))
	fmt.Println("math.Round(a) = ", math.Round(a))
	fmt.Println("math.Pow(c,b) = ", math.Pow(c, b))

}
