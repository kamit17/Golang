package main

import "fmt"

func main() {

	//strings

	//var nameOne string = "mario"
	//var nameTwo = "luigi"
	// declare variable to be used later
	//var nameThree string

	//fmt.Println(nameOne, nameTwo, nameThree)
	// changing value of declared variable
	//nameOne = "peach"
	//nameThree = "bowser"

	//fmt.Println(nameOne, nameTwo, nameThree)

	//using shorthand  to delcare variable for the first time
	//nameFour := "user"
	//fmt.Println(nameFour)

	//integers -- int for whole numbers and float for decimal point

	//ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//bits & memory
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint = 255

	//float

	///var scoreOne float32 = 25.98
	//var scoreTwo float64 = 889899.12321
	//scoreThree := 1.5

	var _firstName string
	var _lastName string
	var _houseNumber string
	var _streetName string
	var _city string
	var _stateAbbreviation string
	var _zipCode string

	fmt.Print("Enter your First Name: ")
	fmt.Scan(&_firstName)

	fmt.Print("Enter your Last Name: ")
	fmt.Scan(&_lastName)

	fmt.Print("Enter your House Number: ")
	fmt.Scan(&_houseNumber)

	fmt.Print("Enter your street name: ")
	fmt.Scan(&_streetName)

	fmt.Print("Enter your City: ")
	fmt.Scan(&_city)

	fmt.Print("Enter your State Abbreviation: ")
	fmt.Scan(&_stateAbbreviation)

	fmt.Print("Enter your zip code: ")
	fmt.Scan(&_zipCode)

	fmt.Println(_firstName + " " + _lastName)
	fmt.Println(_houseNumber + " " + _streetName + " " + "St")
	fmt.Println(_city + " " + _stateAbbreviation + " " + _zipCode)

	/*
		firstname := "Robert"
		streetaddress := "140 highstrung street"
		//streetaddress := "140 highstrung street"
		year_birth := 1974

		var state_of_birth string
		fmt.Print("Enter your state of birth: ")
		fmt.Scanln(&state_of_birth)

		var state_of_residence string
		fmt.Print("Where do you live now? : ")
		fmt.Scanln(&state_of_residence)

		fmt.Println("First name is : ", firstname)
		fmt.Println("streetaddress is: ", streetaddress)
		fmt.Println("Year of birth is: ", year_birth)

		fmt.Println("Your state of birth  is: ", state_of_birth)

		fmt.Println("You now live in: ", state_of_residence)

		// use float64 for most time
	*/
}
