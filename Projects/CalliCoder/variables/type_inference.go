// When you declare a variable with an initial value, golang automatically infers the type of variable from the value on the right-hand side

package main

import "fmt"

func main() {
	var name = "Sanjeev Batra"
	fmt.Printf("Variable 'name' is of type %T\n", name)

}
