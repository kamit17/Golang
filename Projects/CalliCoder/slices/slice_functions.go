package main

import "fmt"

func main() {

	fmt.Println("Example of Copy function")
	src := []string{"Sublime", "VSCode", "IntelliJ", "Eclipse"}
	dest := make([]string, 2)

	numElementsCopied := copy(dest, src)

	fmt.Println("src = ", src)
	fmt.Println("dest = ", dest)
	fmt.Println("Number of elements copied from src to dest = ", numElementsCopied)

	//copy function end

	// append function start
	fmt.Println("Example of append function")
	slice1 := []string{"C", "C++", "Java"}
	slice2 := append(slice1, "Python", "Ruby", "Go")
	fmt.Printf("slice1 = %v, len = %d, cap = %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 = %v, len = %d, cap = %d\n", slice2, len(slice2), cap(slice2))

	slice1[0] = "C#"
	fmt.Println("\nslice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	// append function end
}
