package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in clean up", r)
	}
	wg.Done()

}

func say(s string) {
	defer cleanup()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Oh dear we have 2")
		}
	}
}

func main() {
	defer wg.Add(1)
	go say("hey")
	wg.Add(1)
	go say("There")
	wg.Wait()
}
