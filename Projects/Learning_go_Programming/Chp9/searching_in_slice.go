/*
Go provides a search function that finds an element and returns
its position in a particular integer, float, or string slice. The search
function works well with a sorted slice in the ascending order. If the
element is not found in a particular slice, then it returns an index value
at which that specific element fits in a slice
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{0, 9, 16, 105, 345}
	slice2 := []string{"Akansha", "Mayank", "Ram"}
	slice3 := []float64{0, 0.001, 5.7}
	fmt.Println(sort.SearchInts(slice1, 105))       // 3
	fmt.Println(sort.SearchInts(slice1, 78))        // 3
	fmt.Println(sort.SearchInts(slice1, 0.0))       // 0
	fmt.Println(sort.SearchStrings(slice2, "Ram"))  // 2
	fmt.Println(sort.SearchStrings(slice2, "ram"))  // 3
	fmt.Println(sort.SearchFloat64s(slice3, 0.000)) // 0
	fmt.Println(sort.SearchFloat64s(slice3, 5.7))   // 2
}
