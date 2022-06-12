/*Go provides a data type called bool to store boolean values. It can have two possible values - true and false */

package main

import "fmt"

func main() {
	var truth = 3 <= 5
	var falsehood = 10 != 10

	//short circuiting

	var res1 = 10 > 20 && 5 == 5 //second operatand is not evaluated
	var res2 = 2*2 == 4 || 10%3 == 0

	fmt.Println(truth, falsehood, res1, res2)
}
