package main

import "fmt"

func main() {
	var arrayval [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := arrayval[1:8]
	slice2 := slice1[:6]

	fmt.Println("original array", arrayval)
	fmt.Println("slice from array:", slice1)
	fmt.Println("slice from existing slice", slice2)

	slice1[0] = 178

	fmt.Println("array after modifying slice", arrayval)
	fmt.Println("len of slice1 is ", len(slice1))
	fmt.Println("capacity of slice1 is ", cap(slice1))
	fmt.Println("slice from array after modifying slice:", slice1)
	fmt.Println("slice from existing slice after modifying slice", slice2)

	slice3 := make([]int, 2, 4) // creates an interger type array of size 4, assisngs a reference to slice till 2nd index
	fmt.Println("length of slice3", len(slice3))
	fmt.Println("capacity of slice3", cap(slice3))
	fmt.Println("slice", slice3)

	slice3 = append(slice3, 100)
	fmt.Println("length of slice", len(slice3))
	fmt.Println("capacity of slice", cap(slice3))
	fmt.Println("slice", slice3)

}
