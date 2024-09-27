package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func operation(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}

func circle(r float64) (circumference float64, area float64) {
	circumference = 2 * 3.14 * r
	area = 3.14 * r * r
	return circumference, area
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// to return onlb  a subset of values, use the blank identifier.
	_, c := vals()
	fmt.Println(c)

	add, sub, mul := operation(a, b)
	fmt.Println("addition: ", add)
	fmt.Println("subtraction: ", sub)
	fmt.Println("multiplication: ", mul)

	circumference, area := circle(float64(a))
	fmt.Println("circumference is :", circumference)
	fmt.Println("area: ", area)

}
