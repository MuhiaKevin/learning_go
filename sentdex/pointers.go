package main

import "fmt"

// use a pointer to add the values
func add(num1, num2 int, sum *int) {
	*sum = num1 + num2
}

func main() {
	// a := 12
	// b := &a      // get location of memory address
	// *b = 16      // change the value that the memory holds
	// *b = *b * *b // 16 * 16
	// fmt.Println(a)
	var sum int
	a, b := 12, 36
	// 16 * 16
	add(a, b, &sum)
	fmt.Println(a, b, sum)
}
