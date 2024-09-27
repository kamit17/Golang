package main

import "fmt"

func main() {

	slice1 := []int{1, 2, 3}
	slice2 := []int{10, 20, 30, 54}

	fmt.Println("destination slice if it's lenght is greater than source slice >>>>>>")
	fmt.Println("destinatin slice before copying elements:", slice2)
	n := copy(slice2, slice1)
	fmt.Println("destination slice after copying elements:", slice2)
	fmt.Println("lenght of destinatin slice after copying elements:", len(slice2))

	fmt.Println("number of copied elsements after copying elements:", n)
	// if(len(des_slice) = len(src_slice))
	slice3 := []int{34, 20, 1, 99}
	fmt.Println("destination slice if it's lenght is greater than source slice >>>>>>")
	fmt.Println("destinatin slice before copying elements:", slice3)
	n = copy(slice3, slice1)
	fmt.Println("destination slice after copying elements:", slice3)
	fmt.Println("lenght of destinatin slice after copying elements:", len(slice3))

	fmt.Println("number of copied elsements after copying elements:", n)

	slice5 := []int{}
	fmt.Println("destination slice if it's lenght is greater than source slice >>>>>>")
	fmt.Println("destinatin slice before copying elements:", slice5)
	n = copy(slice5, slice1)
	fmt.Println("destination slice after copying elements:", slice5)
	fmt.Println("lenght of destinatin slice after copying elements:", len(slice5))

	fmt.Println("number of copied elsements after copying elements:", n)
}
