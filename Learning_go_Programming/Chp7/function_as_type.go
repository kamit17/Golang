// syntax : type sub func(a int, b int) int

package main

import "fmt"

type sub func(a int, b int) int

func main() {
	var result sub = func(a int, b int) int {
		return a - b
	}
	fmt.Println("Result is: ", result(100, 70))
}
