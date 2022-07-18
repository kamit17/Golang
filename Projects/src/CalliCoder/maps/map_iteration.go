//map can be iterated using range to give key value pair inevery iteration

package main

import "fmt"

func main() {
	var personAge = map[string]int{
		"Rajeev": 25,
		"James":  32,
		"Sarah":  29,
	}

	for name, age := range personAge {
		fmt.Println(name, age)
	}
}
