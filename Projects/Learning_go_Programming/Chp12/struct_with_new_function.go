package main

import "fmt"

func main() {
	type student struct {
		name,city  string
		rollno int			
	}
	p1 := new(student)
	fmt.Println(*p1)
	p1.city='jaipur'

	p1.rollno =67
	fmt.Println(*p1)
	//p1.grade= 'Pass'
	
	//anonnymous struct with new

	anony_point :=new(struct{
		name,city string
		//grade string
		rollno int
	})
	fmt.Println(*anony_point)
	anony_point.city = 'Delhi'
	anony_point.name = 'Abhish'
	fmt.Println(*anony_point)

}