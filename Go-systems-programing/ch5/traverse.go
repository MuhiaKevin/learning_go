package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunction(path string, info os.FileInfo, err error) error {
	_, err = os.Stat(path)

	if err != nil {
		return err
	}

	fmt.Println(path)

	return nil
}

func main() {
	argurements := os.Args

	if len(argurements) == 1 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}
	Path := argurements[1]
	err := filepath.Walk(Path, walkFunction)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
