package main

import (
	"fmt"
	"time"
)

func printHello() {
	// fmt.Println("Started printHello Scheduler")
	// time.Sleep(20000 * time.Millisecond)
	fmt.Println("Hello world")
}

func main() {
	fmt.Println("main exectution started ")
	// create a goroutine
	go printHello()
	// schedule another goroutine
	time.Sleep(10 * time.Millisecond)
	fmt.Println("main execution stopped")
}
