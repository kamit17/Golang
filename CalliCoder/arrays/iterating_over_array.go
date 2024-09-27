package main

import "fmt"

func main() {
	brands := [3]string{"Fossil", "Kirkland", "logitech"}

	for i := 0; i < len(brands); i++ {
		fmt.Println(brands[i])
	}
}
