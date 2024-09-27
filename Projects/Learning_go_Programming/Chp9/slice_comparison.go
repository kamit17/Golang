package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice1 := []int{23, 0, 105, 0}
	slice2 := []string{"Ram", "Akansha", "Mayank"}
	slice3 := []float64{23, 0, 105, 0}
	slice4 := make([]int, 4)
	slice4[0] = 23
	slice4[1] = 0
	slice4[2] = 105
	fmt.Println("Is slice1 equal to slice2:", reflect.DeepEqual(slice1, slice2))
	fmt.Println("Is slice1 equal to slice3:", reflect.DeepEqual(slice1, slice3))
	fmt.Println("Is slice1 equal to slice4:", reflect.DeepEqual(slice1, slice4))
}
