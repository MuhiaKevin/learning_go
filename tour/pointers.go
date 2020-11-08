package main

import "fmt"

func main() {
	var myName string
	// a pointer to an interger
	var p *string
	myName = "Kevin Muhia"
	//
	p = &myName
	fmt.Println(*p) // read through the pointer p
	fmt.Println(p)  // print the memeory location of the varaible
}
