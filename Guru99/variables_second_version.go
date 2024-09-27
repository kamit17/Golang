package main

import "fmt"

func main() {
	//declaring an integer variable x
	var x int
	x = 3                //assigning x the value 3
	fmt.Println("x:", x) //prints3

	//declaring an integer variable y with value 20 in  a single statement andprint it
	var y int = 20
	fmt.Println("y:", y)

	//declare variable z with value 5 and print it
	//Here type int is not explicitly mentioned
	var z = 50
	fmt.Println("z", z)

	//Multiple variables are assigned in single lin -i with an interger and j with a string

	var i, j = 100, "hello"
	fmt.Println("i and j:", i, j)

	//declare and assing a variable using :
	a := 20
	fmt.Println(a)

	//gives a error since a is already declared
	a := 30
	fmt.Print(30)
}
