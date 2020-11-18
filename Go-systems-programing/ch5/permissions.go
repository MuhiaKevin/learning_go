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

	fileinfo, err := os.Stat(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	mode := fileinfo.Mode()
	fmt.Print(filename, " : ", mode, "\n")
}
