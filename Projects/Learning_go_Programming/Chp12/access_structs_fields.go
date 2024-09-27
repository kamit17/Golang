package main

import (
	"fmt"
)

func main() {
	var student struct {
		name, city string
		rollno     int
	}

	student.name = "William"
	student.city = "Delhi"
	student.rollno = 101

	fmt.Println(student)

	type student1 struct {
		name, city string
		rollno     int
	}
	var student_var student1
	student_var.name = "John"
	student_var.city = "Delhi"
	fmt.Println(student_var)
}
