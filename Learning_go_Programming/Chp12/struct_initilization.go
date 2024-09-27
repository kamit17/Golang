package main

import "fmt"

func main() {
	type student struct {
		name, city string
		rollno     int
	}

	var student_var student

	student_var = student{rollno: 67, city: "jaipur"}
	fmt.Println(student_var)

	//Anonymous struct

	student1 := struct {
		name   string
		city   string
		rollno int
	}{
		name: "John",
		city: "Sanjose",
	}
	fmt.Println(student1)
}
