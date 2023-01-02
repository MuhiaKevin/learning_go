package main

import "fmt"

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

// channeles are used together with goroutines
func main() {
	fooVal := make(chan int)
	// fmt.Println(fooVal)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	v1 := <-fooVal
	// v2 := <-fooVal
	fmt.Println(v1)

}
