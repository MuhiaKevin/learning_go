package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // point to i
	k := &j // point to i
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i

	fmt.Println(*p, *k) // read i through the pointer
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j
}
