//takes two integer input values and returns the sum of them

package main

import "fmt"

func sum(a int, b int) int {
	sum_value := a + b
	fmt.Println("sum is:", sum_value)
	return sum_value

}

func main() {
	a := 90
	b := 30
	sum(a, b)
}
