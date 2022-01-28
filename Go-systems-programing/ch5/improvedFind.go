package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// used to exclude a particular file
func excludeNames(filename string, exclude string) bool {
	if exclude == "" {
		return false
	}
	if filepath.Base(filename) == exclude {
		return true
	}
	return false
}

func main() {
	// flags to tell program what we want to be printed
	minusS := flag.Bool("s", false, "Sockets")
	minusP := flag.Bool("p", false, "Pipes")
	minusSL := flag.Bool("sl", false, "Symbolic Links")
	minusD := flag.Bool("d", false, "Directories")
	minusF := flag.Bool("f", false, "Files")
	minusX := flag.String("x", "", "Files")

	flag.Parse()
	flags := flag.Args()

	printAll := false

	// if all the flags are unset, the program will print everything
	if *minusS && *minusP && *minusSL && *minusD && *minusF {
		printAll = true
	}
	//if all the flags are set, the program will also print everything
	if !(*minusS || *minusP || *minusSL || *minusD || *minusF) {
		printAll = true
	}

	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	Path := flags[0]

	walkFunction := func(path string, info os.FileInfo, err error) error {
		fileInfo, err := os.Stat(path)

		if err != nil {
			return err
		}

		if excludeNames(path, *minusX) {
			return nil
		}

		if printAll {
			fmt.Println(path)
			return nil
		}
		mode := fileInfo.Mode()

		if excludeExtensions(path, *minusEXT))

		if mode.IsRegular() && *minusF {
			fmt.Println(path)
			return nil
		}

		if mode.IsDir() && *minusD {
			fmt.Println(path)
		}

		fileInfo, _ = os.Stat(path)
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			if *minusSL {
				fmt.Println(path)
				return nil
			}
		}
		if fileInfo.Mode()&os.ModeNamedPipe != 0 {
			if *minusP {
				fmt.Println(path)
				return nil
			}
		}

		if fileInfo.Mode()&os.ModeSocket != 0 {
			if *minusS {
				fmt.Println(path)
				return nil
			}
		}
		return nil
	}

	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
