// a zero value of a pointer is null. i.e any unitizalized pointer will have value nil

package main

import "fmt"

func main() {
	var p *int
	fmt.Println("p = ", p)
}
