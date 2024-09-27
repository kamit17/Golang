package main

import "fmt"

func main() {
	var age int8 = 32           //signed 8 bit integer
	var port int16 = 80         //signed 16 bit integer
	var zipcode int32 = 90348   // signed 32 bit
	var phone int64 = 732335624 //signed 64 bit

	var phone2 uint64 = 7323356354 //unsinged 64 bit
	var score int64 = -1           //unsinged 64 bit with -ve value

	fmt.Println("age int8", age)
	fmt.Println("port int16", port)
	fmt.Println("zipcode int32", zipcode)
	fmt.Println("phone int64", phone)
	fmt.Println("phone2 uint64", phone2)
	fmt.Println("score int64", score)

}
