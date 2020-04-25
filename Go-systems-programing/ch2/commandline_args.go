package main

import (
	"fmt"
	"os"
)

// Running a go program :-> go run hw.go
// go build hw.go -> builds an executeable which you run as ./hw

func main() {
	arguments := os.Args
	names := ""

	for i := 0; i < len(arguments); i++ {
		// fmt.Println(arguments[i])
		names += arguments[i]
		names += " "
	}

	fmt.Println(names)
}
