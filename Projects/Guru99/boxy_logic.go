package main

import (
	"fmt"
	"strconv"
)

func main() {
	var _length string
	var _width string
	var _height string

	var _Length int
	var _Width int
	var _Height int

	var _area int

	fmt.Print("Enter the width: ")
	fmt.Scan(&_width)
	fmt.Print("Enter the length: ")
	fmt.Scan(&_length)
	fmt.Print("Enter the height: ")
	fmt.Scan(&_height)

	//conver string to int

	_Length, _ = strconv.Atoi(_length)
	_Width, _ = strconv.Atoi(_width)
	_Height, _ = strconv.Atoi(_height)

	_area = _Height * _Length * _Width
	fmt.Println("Area is : ", _area)
}
