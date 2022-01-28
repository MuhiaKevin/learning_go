package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunction(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return err
	}

	mode := fileInfo.Mode()
	// this will ensure that the path is a directory or a file and not a symbolic link or a named pipe
	if mode.IsDir() || mode.IsRegular() {
		fmt.Println(path)
	}
	return nil
}

func main() {
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	Path := flags[0]
	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
