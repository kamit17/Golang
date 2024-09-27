package main

import "fmt"

func main() {
	type student struct {
		name, cit string
		rollno    int
	}

	var student_var student
	fmt.Println(student_var)

	var student1 struct {
		name, city string
		rollno     int
	}

	fmt.Println(student1)

}
