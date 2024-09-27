package main

import (
	"fmt"
)

func main() {
	s1 := [][]int{
		{11, 22},
		{33, 44},
		{55, 66},
	}
	fmt.Println(s1[2][0])

	s2 := [][]string{
		[]string{"golang", "python", "MySQL"},
		[]string{"HTML", "CSS", "JavaScript"},
	}

	fmt.Println(s2)
}
