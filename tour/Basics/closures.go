package main

import "fmt"

/*
https://www.calhoun.io/what-is-a-closure/

A closure is a function value that references variables from outside its body.
A closure is a special type of anonymous function that references variables declared outside of the function itself
The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
*/

// anonymous function
var DoStuff func() = func() {
	fmt.Println("Hello world")
}

// RegFunc
func RegFunc() { fmt.Println("reg func") }

func newCounter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
func main() {

	// TODO: 1. ANONYMOUS FUNCTIONS

	// overwrite the function
	DoStuff = func() {
		fmt.Println("Hello Africa")
	}
	// assign another function to a anonymous function
	DoStuff = RegFunc
	DoStuff()

	// TODO: 2. ANONYMOUS FUNCTIONS
	n := 0
	counter := func() int {
		n++
		return n

	}
	// anonymous function has access to the n variable, but this was never passed in as a parameter when counter() was called. This is what makes it a closure!
	fmt.Println(counter()) // n=1
	fmt.Println(counter()) // n=2

	counter1 := newCounter()
	fmt.Println(counter1())
	fmt.Println(counter1())
}
