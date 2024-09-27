package main

import "fmt"

func main() {

	a := 5
	b := 10
	c := 4
	d := 30
	e := 89
	f := 1

	// Statements that evaluate to True
	fmt.Printf("a < b = %v\n", a < b)   // True: 5 is less than 10
	fmt.Printf("a == c = %v\n", a == c) // True: 5 is equal to 5
	fmt.Printf("b > d = %v\n", b > d)   // False: 10 is not greater than 20

	// Statements that evaluate to False
	fmt.Printf("a > b = %v\n", a > b)   // False: 5 is not greater than 10
	fmt.Printf("b == d = %v\n", b == d) // False: 10 is not equal to 20
	fmt.Printf("e != f = %v\n", e != f) // True: 0 is not equal to 1

}
