// package declaration
package main

//Importing packages
import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	//Finding the max of two numbers
	fmt.Println(math.Max(73.15, 92.56))

	//Calculate the square root of a number
	fmt.Println(math.Sqrt(225))

	//printing the value of pi
	fmt.Println(math.Pi)

	// Epoch time in milliseconds
	epoch := time.Now().Unix()
	fmt.Println(epoch)

	// Generating a random integer between 0 to 100
	rand.Seed(epoch)
	fmt.Println(rand.Intn(100))

}
