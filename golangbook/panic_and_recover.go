package main

import "fmt"

func defFoo() {
	fmt.Println("defFoo() started")

	if r := recover(); r != nil {
		fmt.Println("WOHA! Program is panicking with value", r)
	}

	fmt.Println("defFoo() done")
}

func normMain() {

	fmt.Println("normMain() started")

	defer defFoo() // defer defFoo call

	panic("HELP") // panic here

	fmt.Println("normMain() done")
}

func defMain() {

	fmt.Println("defMain() started")

	fmt.Println("defMain() done")
}

func main() {
	fmt.Println("main() started")

	defer defMain() // defer call

	normMain() // normal call

	fmt.Println("main() done")
}

// package main

// import "fmt"

// func main() {
//   defer func() {
//     str := recover()
//     fmt.Println(str)
//   }()
//   panic("PANIC")
// }
