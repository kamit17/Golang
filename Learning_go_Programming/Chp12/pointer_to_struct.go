package main

import (
	"fmt"
)

func main() {
	type student struct {
		name, city string
		rollno     int
	}

	student_var := student{"Jim", "Delhi", 1}
	fmt.Println("Creating a struct  and assign its address to a pointer")
	p1 := &student_var
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(p1.name)
	fmt.Println(p1.city)
	fmt.Println((*p1).city)
	fmt.Println(p1.rollno)

	//DIrectly create a pointer

	fmt.Println("Directly created pointer")
	p2 := &student{"jim", "Delhi", 67}
	fmt.Println(p2)
	fmt.Println(*p2)
	fmt.Println(p2.name)
	fmt.Println(p2.city)
	fmt.Println((*p2).city)

	// pointer with anonymous struct

	fmt.Println("Pointer with Anonymous struct")
	var student1 struct {
		name, city string
		rollno     int
	}

	annoy_point := &student1
	annoy_point.name = "Jim"
	annoy_point.rollno = 67

	fmt.Println(annoy_point)
	fmt.Println(*annoy_point)
	fmt.Println(annoy_point.name)
	fmt.Println(annoy_point.rollno)

}
