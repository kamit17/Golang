//You can add new items to an initialized map using the following syntax
//m[key] = value

//initializes a map using make() funcaiton and add some items to it

package main

import "fmt"

func main() {
	//initializing a map
	var tinderMatch = make(map[string]string)

	//Adding keys to a map
	tinderMatch["James"] = "Zoey" // Assigns  value Zoey to James
	tinderMatch["Dave"] = "Sophie"
	tinderMatch["Jim"] = "Katy"

	fmt.Println(tinderMatch)

	/*
		Adding a key that already exists will override the existing key with a new value
	*/
	tinderMatch["Jim"] = "Pam"
	fmt.Println((tinderMatch))
}
