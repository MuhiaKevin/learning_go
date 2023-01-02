package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := z; i < x; i++ {
		if z != math.Sqrt(x) {
			// fmt.Println(z)
			z -= (z*z - x) / (2 * z)
		} else {
			return z
			break
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(9))

}
