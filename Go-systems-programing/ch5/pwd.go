package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	argurements := os.Args
	pwd, err := os.Getwd()

	if err == nil {
		fmt.Println(pwd)
	} else {
		fmt.Println(err)
	}

	if len(argurements) == 1 {
		return
	}

	if argurements[1] != "-P" {
		return
	}
	fileinfo, err := os.Lstat(pwd)
	if fileinfo.Mode()&os.ModeSymlink != 0 {
		fmt.Println(pwd, "is a symbolic link")
		realpath, err := filepath.EvalSymlinks(pwd)

		if err == nil {
			fmt.Println("Path: ", realpath)
		}
	}

}
