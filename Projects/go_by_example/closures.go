package main

import "fmt"

//function intSeq returns another function which we define anonymously in the body of intSeq.
//THe returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// we call intseq , assigning the result to nextINt. This function value vaptures its own i value which will be updated each time we call nextInt.
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	//To confirm that the state is unique to that particular function, create and test a new one.

	newInts := intSeq()
	fmt.Println(newInts())
}
