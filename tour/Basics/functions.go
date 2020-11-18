package main

import (
	"fmt"
)

// https://medium.com/rungo

// function to add two numbers and return a value
// params {a,b of type int} return value an integer
func add(num1, num2 int) int {
	return num1 + num2
}

// functions can return more than one values
func square(num1 int) (int, string) {
	if num1 == 0 {
		return 0, "Value should greater than 0"
	} else {
		return (num1 * num1), ""
	}
}

// function can also return named values
func subtract(num1, num2 int) (result int) {
	result = num1 - num2
	return
}

// functions can also get functions as parameters

func calculate(a int, b int, aFunction func(int, int) int) int {
	result := aFunction(a, b)
	return result
}

// function main is where the programs starts
func main() {
	// in fmt.Println P is capital is in caps because it is exported
	// fmt.Println("This is a function")

	// square, error := square(0)

	// if len(error) > 0 {
	// 	fmt.Println(error)
	// 	return
	// }

	// fmt.Println(square)

	// fmt.Println(subtract(10, 5))
	addResult := calculate(25, 25, add)
	subResult := calculate(100, 25, subtract)
	fmt.Printf("Addition value is :  %v Subtraction value is %#v\n", addResult, subResult)

}
