package main

import (
	"fmt"
	"runtime"
)

func showOperatingSystem() {
	os := runtime.GOOS
	switch os {
	case "linux":
		fmt.Println("Go is running on a linux maachine")
	case "darwin":
		fmt.Println("Go is running on a linux maachine")
	// print windows and etc
	default:
		fmt.Printf("Go is running on %s \n", os)
	}
}

func sayGoodBye() {
	fmt.Printf("This function will run last \n")
}
func sayHello() {
	fmt.Printf("Heey people \n")
}

// go uses a stack data structure to call  functions that have a defer keyword
func main() {
	defer sayHello()
	defer sayGoodBye()
	showOperatingSystem()
}
