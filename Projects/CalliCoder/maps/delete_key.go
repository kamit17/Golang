//Delete the key from the map

package main

import "fmt"

func main() {
	var fileExtensions = map[string]string{
		"Python": ".py",
		"C++":    ".cpp",
		"Java":   ".java",
		"Golang": ".go",
		"Kotlin": ".kt",
	}

	fmt.Println(fileExtensions)

	delete(fileExtensions, "Kotlin")

	//delete function does do anything ig the key doesnit exist
	delete(fileExtensions, "Javascript")

	fmt.Println(fileExtensions)
}
