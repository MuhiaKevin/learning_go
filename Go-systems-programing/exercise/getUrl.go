package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("usage: %s message  filename\n", filepath.Base(arguments[]))
	}
}
