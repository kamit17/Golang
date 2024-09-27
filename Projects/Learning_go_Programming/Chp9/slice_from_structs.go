package main

import "fmt"

func main() {
	record := []struct {
		rollno  int
		name    string
		present bool
	}{
		{1, "A", true},
		{2, "B", false},
		{3, "C", true},
		{4, "D", false},
	}
	fmt.Println("structs type slice is :", record)
}
