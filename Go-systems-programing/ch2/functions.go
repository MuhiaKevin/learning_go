package main

import (
	"fmt"
)

// set the function parameters and the data type of the return values

func unNamedMinMax(x, y int) (int, int) {
	if x > y {
		min := y
		max := x
		return min, max
	} else {
		min := x
		max := y
		return min, max
	}
}

// set the function parameters and the data type of the return values

func minmax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
		return min, max
	} else {
		min = x
		max = y
		return min, max
	}
}

func main() {
	// fmt.Println("Answer from unNamedMinmax()", unNamedMinMax(235, 110))
	fmt.Println(minmax(235, 110))
}
