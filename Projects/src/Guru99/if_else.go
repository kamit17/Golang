package main

import "fmt"

func main() {
	var x = 50
	if x < 10 {
		//Executes if x is less than 10
		fmt.Println("x is less than 10")

	} else if x >= 10 && x <= 90 {
		fmt.Println("x is between 10 and 90")
	} else {
		fmt.Println("x is greater  than 90")
	}

}
