package main

// the package to be created
import "calculation"

func main() {
	x, y := 15, 10
	//the package will have functin Do_add()
	sum := calculation.Do_add(x, y)
}
