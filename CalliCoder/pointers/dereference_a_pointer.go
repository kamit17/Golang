// * can be usd to access the value store din the variable that the pointer points to

package main

import "fmt"

func main() {

	var a = 100
	var p = &a

	fmt.Println("a =  ", a)
	fmt.Println("p = ", p)
	fmt.Println("*p = ", *p)
}
