package main

import (
	"fmt"
	"math"
)

// functions can take functions as values
func compute(fn func(float64, float64) float64) float64 {
	return fn(4, 3)
}

func main() {
	hypot := func(x, y float64) float64 {
		return (math.Sqrt(x*x + y*y))
	}

	fmt.Println(hypot(4.0, 9.0))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

}
