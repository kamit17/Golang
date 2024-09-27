package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Using a fixed seed for reproducibility
	fixedSeed := int64(42)
	src1 := rand.NewSource(fixedSeed)
	r1 := rand.New(src1)

	fmt.Println("Random numbers with fixed seed:")
	for i := 0; i < 5; i++ {
		fmt.Println(r1.Intn(100)) // Generates a random number between 0 and 99
	}

	// Using the current time as a seed for true randomness
	src2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(src2)

	fmt.Println("\nRandom numbers with current time seed:")
	for i := 0; i < 5; i++ {
		fmt.Println(r2.Intn(100)) // Generates a random number between 0 and 99
	}
}
