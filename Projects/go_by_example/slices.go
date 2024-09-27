package main

import (
	"fmt"
)

func main() {
	/*var s []string                                   // slices are typed only by the elements they contain
	fmt.Println("uninit:", s, s == nil, len(s) == 0) // Unintialzed slice equals to nil and has length 0

	s = make([]string, 3)                                  // creating a slice of strings of length 3 using builtin make
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // if we know that a new slice is going to grow ahead o time, it's possible to pass a capacity explicity as an additional parameter to make.
	/*
		s[0] = "a"
		s[1] = "b"
		s[2] = "c"
		fmt.Println("set:", s)
		fmt.Println("get:", s[2])

		fmt.Println("len:", len(s))
		//appending to a slice
		s = append(s, "d")
		s = append(s, "e", "f")
		fmt.Println("apd:", s)

		// copying a slice
		c := make([]string, len(s))
		copy(c, s)
		fmt.Println("cpy", c)

		// using slice operator with syntax slice[low:high] to get a slice
		l := s[2:5]
		fmt.Println("sl1:", l) // gets a slice of s[2],s[3],s[4]

		l = s[:5] // slices upto (but exlcuging) s[5]
		fmt.Println("sl2:", l)

		l = s[2:] // slice up from and including s[2]
		fmt.Println("sl3:", l)

		t := []string{"g", "h", "i"} // declaring and intializing a slice in a single line
		fmt.Println("dcl:", t)

		t2 := []string{"g", "h", "i"}
		if slices.Equal(t, t2) {
			fmt.Println("t == t2")
		}
	*/

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Println("i: ", i)
		innerLen := i + 1
		fmt.Println("inner length: ", innerLen)
		twoD[i] = make([]int, innerLen)
		fmt.Println("initial twoD: ", twoD[i])
		for j := 0; j < innerLen; j++ {
			fmt.Println("j: ", j)
			twoD[i][j] = i + j
			fmt.Println("inside loop 2d: ", twoD)

		}
		fmt.Println("outside loop 2d: ", twoD)
	}
	fmt.Println("2d: ", twoD)
}
