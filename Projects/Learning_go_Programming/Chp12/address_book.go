//Golang program to show how to declare and define structs

package main

import "fmt"

//Defining a struct type
type Address struct {
	Name, city, State string
	zipcode           int
}

func main() {
	//  Declaring a variable of a `struct` type
	// All the struct fields are initialized
	// with their zero value
	var a Address
	fmt.Println(a)

	//Declaring and initilizing a struct using struct literal
	a1 := Address{"Akshay", "Dehradun", "Uttarakhand", 248}

	fmt.Println("Address: ", a1)

	// Naming fields while
	// initializing a struct
	a2 := Address{Name: "Anikaa", city: "Ballia",
		zipcode: 277001}

	fmt.Println("Address2: ", a2)

	// Uninitialized fields are set to
	// their corresponding zero-value
	a3 := Address{Name: "Delhi"}
	fmt.Println("Address3: ", a3)

}
