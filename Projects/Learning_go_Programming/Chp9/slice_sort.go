package main

import (
	"fmt"
	"sort"
)
func main() {
	slice1 := [] int{23, 0, 105, 58}
	slice2 := [] string{"Ram", "Akansha", "Mayank"}
	slice3 := [] float64{75.84, 0.007, 85.62, 0.74637}
	fmt.Println("slice1 before sorting:", slice1)
	sort.Ints(slice1)
	fmt.Println("slice1 after sorting:", slice1)
	fmt.Println("slice1 before sorting:", slice2)
	sort.Strings(slice2)
	fmt.Println("slice1 after sorting:", slice2)
	fmt.Println("slice1 before sorting:", slice3)
	sort.Float64s(slice3)
	fmt.Println("slice1 after sorting:", slice3)
}