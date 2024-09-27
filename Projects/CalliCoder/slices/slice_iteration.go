package main

import "fmt"

func main() {
	countries := []string{"India", "USA", "England", "Germany"}

	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}
	// iteration using range form of loop
	primeNumbers := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	for index, number := range primeNumbers {
		fmt.Printf("PrimeNumber(%d) = %d\n", index+1, number)
	}
	// to ignore the index  we can use _

	numbers := []float64{3, 5, 7, 2, 3, 9, 9.2, 12.1}

	sum := 0.0

	for index, number := range numbers {
		sum += number
	}
	fmt.Printf("Total Sum = %.2f\n", sum)
}
