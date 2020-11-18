package main

import (
	"fmt"
	"math"
)

// this is like a class Coordinates in js or python
// Coordinates
type Coordinates struct {
	Lat  float64
	Long float64
}

// methods can also be written as below
func Abs(v Coordinates) float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

// this is like a method of a class in js or python
// Coordinates
func (v Coordinates) Abs() float64 {
	return math.Sqrt(v.Lat*v.Lat + v.Long*v.Long)
}

func main() {
	vertex := Coordinates{3, 4}
	fmt.Println(vertex)
	// call the Abs() method
	fmt.Println(vertex.Abs())
	// this is another way of creating and calling a method in golang
	fmt.Println(Abs(vertex))

}
