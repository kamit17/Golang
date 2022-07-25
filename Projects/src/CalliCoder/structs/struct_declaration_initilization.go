package main

import (
	"fmt"
)

//Defining a struct

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	//Declaring a variable of a struct type
	var p Person // All strunct fileds are initialized with their zero value
	fmt.Println(p)

	//Declaring and initilizing a strunt using a struct literal

	p1 := Person{"Rajeev", "Singh", 26}
	fmt.Println("Person1: ", p1)

	//Naming field while initilizing a struct

	p2 := Person{
		FirstName: "John",
		LastName:  "Snow",
		Age:       45,
	}
	fmt.Println("Person2: ", p2)

	// Uninitialized fields are set to their corresponding zero-value
	p3 := Person{FirstName: "Robert"}
	fmt.Println("Person3: ", p3)
}
