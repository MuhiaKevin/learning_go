// methods with pointer receivers change the values of struct
package main

import (
	"math"
)

// Vertex
type Vertex struct {
	X, Y float64
}

// method Vertex
func (v Vertex) Abs() float64 {
	v.X = 12
	v.Y = 12

	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Vertex pointer Receiver
// it will change a copy of the  values
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Vertex Pointer Receiver Scale as a function
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	vertex := Vertex{3, 4}
	vertex.Scale(16)
	Scale(&vertex, 10) // rewritten as a function
	// fmt.Println(vertex.Abs())
	// fmt.Println(vertex.Abs())
}
