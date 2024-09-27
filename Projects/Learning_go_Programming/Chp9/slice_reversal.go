/*
The reverse function returns a slice in the reverse/descending order. It
works well with sorted or unsorted slices.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{0, 888, 9, 16, 105, 345}
	slice2 := []string{"Soni", "Akansha", "Charu", "Mayank", "Ram"}
	sort.Sort(sort.Reverse(sort.IntSlice(slice1)))
	fmt.Println("reverse of slice1:", slice1)
	sort.Sort(sort.Reverse(sort.StringSlice(slice2)))
	fmt.Println("reverse of slice2:", slice2)
}
