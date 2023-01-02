package main

import "fmt"

func main() {
	var c chan int       //zero-value of an channel is nil
	var i int            // zero-value of an interger is zero
	c2 := make(chan int) // this is used instead so as to make a ready to use channel
	fmt.Println(c, i)    // returns nil because the zero value of a channel is zero
	fmt.Printf("type of `c` is %T\n", c2)
	fmt.Printf("value of `c` is %v\n", c2) // value is a pointer. channels by default are pointers
}
