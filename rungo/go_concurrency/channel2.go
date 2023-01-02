package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

// func greet(c chan string) {
// 	fmt.Println("Hello " + <-c + "!")
// }

// func sendHello(c chan string) {
// 	c <- "Muhia"
// }
func mutiply(c chan int) {
	num := 13
	for {
		if num > 2500000000 {
			c <- num
			break
		}
		num *= 28
	}
}

// func mutiply() int {
// 	var num int = 13
// 	for {
// 		if num > 2500000000 {
// 			return num
// 		}
// 		num *= 28
// 	}
// }

func greet(c chan int) {
	fmt.Println("", <-c)
}

// func greet(x int) {
// 	fmt.Println("", x)
// }
func main() {
	// fmt.Println("main() started")
	// c := make(chan string)
	// go greet(c)
	// go sendHello(c)
	// time.Sleep(10 * time.Millisecond)
	// // c <- "john"
	// fmt.Println("main() stopped")

	c := make(chan int)
	go mutiply(c)
	greet(c)
	fmt.Println("\nmain execution stopped at time", time.Since(start))

	// fmt.Println("\nmain execution stopped at time", time.Since(start))
}
