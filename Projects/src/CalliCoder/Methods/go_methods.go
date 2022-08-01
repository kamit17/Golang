// A method is nothing but a functin with a special ceiver argument

//func (receiver Type) MethodName(paramenter list) (returnTypes){

//}

//Define a method and invoke it

package main

import "fmt"

//Strunct type - 'point'

type Point struct {
	X, Y float64
}

//Method with receiver 'Point'

func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}

func main() {
	p := Point{2.0, 4.0}

	fmt.Println("Point : ", p)

	fmt.Println("Is Point p located above the line y = 1.0 ? : ", p.IsAbove(1))
}
