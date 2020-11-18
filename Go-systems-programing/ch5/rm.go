package main

import (
	"fmt"
	"os"
)

func main() {
	argurements := os.Args
	if len(argurements) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}
	filename := os.Args[1]
	error := os.Remove(filename)

	if error != nil {
		fmt.Println("Error:", error)
		os.Exit(1)
	}
	fmt.Println(filename, "has been deleted")
}
