// short declaration is done using := operator

package main

import "fmt"

func main() {
	name := "Mello Jello"
	age, salary, isProgrammer := 35, 50000.0, true

	fmt.Println(name, age, salary, isProgrammer)
}
