package main

import "fmt"

//globala variable

//var message string = "Bye world!"

func main() {
	fmt.Println("Enter your first name: ")

	var firstName string
	fmt.Scanln(&firstName)

	fmt.Println("Enter your last name: ")
	var secondName string
	fmt.Scanln(&secondName)

	fmt.Println("Your first name is : ")
	fmt.Println(firstName)

	fmt.Println("Your second name is :  ")
	fmt.Println(secondName)

	fmt.Println("Your full name is : ", firstName+" "+secondName)
	//fmt.Println(firstName + " " + secondName)

}
