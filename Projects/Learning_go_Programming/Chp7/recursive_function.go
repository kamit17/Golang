//recursive function can be defined as a function that calls itself directly or indirectly a finit number of times untill the recursive condition becomes false

package main

import "fmt"

func factorial(num int) int {
	if num > 1 {
		return num * factorial(num-1)
	} else {
		return 1
	}
}

func main() {
	num := 10
	result := factorial(num)
	fmt.Println("factorial result is : ", result)
}
