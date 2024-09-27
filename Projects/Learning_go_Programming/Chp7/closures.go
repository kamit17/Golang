package main

import "fmt"

func return_anon() func() int { // function return_anon returns an anonymous fnction
	count := 0
	return func() int {
		count++
		return count
	}
}
func main() {
	return_anon_variable := return_anon()                                    //return_anon_variable is a function returned by return_anon
	fmt.Println("value returned by 'return_anon' is:", return_anon_variable) //prints ‘ox47da70’
	fmt.Println("count is:", return_anon_variable())                         //prints“count is 1”
	fmt.Println("count is:", return_anon_variable())                         //prints “count is 2”
	return_anon_variable2 := return_anon()                                   //new instance
	fmt.Println("count is:", return_anon_variable2())                        //prints“count is 1”
}
