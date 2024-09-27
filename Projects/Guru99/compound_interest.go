package main

import (
	"fmt"
	"math"
)

func main() {
	// Variables to hold user inputs
	var initialDeposit float64
	var interestRate float64
	var timesPerYear int
	var years int

	// Prompt the user for each value
	fmt.Print("Enter the initial deposit amount: ")
	fmt.Scan(&initialDeposit)

	fmt.Print("Enter the annual interest rate (as a decimal, e.g., 0.05 for 5%): ")
	fmt.Scan(&interestRate)

	fmt.Print("Enter the number of times interest is calculated per year: ")
	fmt.Scan(&timesPerYear)

	fmt.Print("Enter the number of years since the initial deposit: ")
	fmt.Scan(&years)

	// Calculate the current value of the deposit using the compound interest formula
	V := initialDeposit * math.Pow(1+interestRate/float64(timesPerYear), float64(timesPerYear*years))

	// Display the results
	fmt.Printf("\nResults:\n")
	fmt.Printf("Initial deposit: $%.2f\n", initialDeposit)
	fmt.Printf("Annual interest rate: %.2f%%\n", interestRate*100)
	fmt.Printf("Number of times interest is calculated per year: %d\n", timesPerYear)
	fmt.Printf("Number of years: %d\n", years)
	fmt.Printf("Current value of the deposit: $%.2f\n", V)
}
