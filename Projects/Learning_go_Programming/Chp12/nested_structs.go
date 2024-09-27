package main

import "fmt"

func main() {
	type result struct {
		math    int
		science int
		english int
		isPass  bool
	}
	type student struct {
		name, city string
		rollno     int
		result
	}
	a := student{"rahul", "jaipur", 11, result{45, 78, 90, true}}
	fmt.Println(a) // {rahul jaipur 11 {45 78 90 true}}
	//anonymous nested struct
	result_annon := struct {
		math    int
		science int
		english int
		isPass  bool
	}{
		math:    45,
		science: 78,
		english: 90,
		isPass:  true,
	}
	student_annon := struct {
		name, city string
		rollno     int
		result
	}{
		name:   "rahul",
		city:   "city",
		rollno: 11,
		result: result_annon,
	}
	fmt.Println(student_annon) // {rahul city 11 {45 78 90 true}}
}
