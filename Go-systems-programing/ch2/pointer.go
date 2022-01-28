package main

import (
	"fmt"
)

// & Ampersand symbols represents the location of a data

func changeValueOfX(x *int) {
	var answer int
	fmt.Print("Enter something other than 0 ")
	fmt.Scanln(&answer)
	// fmt.Println(answer)
	*x = answer
}

func main() {
	var answer int
	fmt.Println(&answer)
	// changeValueOfX(&answer)
	// fmt.Println("New Value of x is : ", answer)

}
