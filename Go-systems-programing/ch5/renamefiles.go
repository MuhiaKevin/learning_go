package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	minusOverwrite := flag.Bool("overwrite", false, "overwrite")
	flag.Parse()
	flags := flag.Args()
	if len(flags) < 2 {
		fmt.Println("Please provide two arguments!")
		os.Exit(1)
	}

	source := flags[0]
	destination := flags[1]

	// to check that source is a  file and not directory
	fileinfo, err := os.Stat(source)
	if err == nil {
		mode := fileinfo.Mode()

		if mode.IsRegular() == false {
			fmt.Println("Sorry, we only support regular files as source!")
			os.Exit(1)
		}
	} else {
		fmt.Println("Error reading:", source)
		os.Exit(1)
	}
	// make sure destination is formated well for example if the destination is a relative path
	newDestination := destination
	destInfo, err := os.Stat(destination)

	if err == nil {
		mode := destInfo.Mode()

		if mode.IsDir() {
			justTheName := filepath.Base(source)
			newDestination = destination + "/" + justTheName
		}
	}

	// check if the destionation file exists
	destination = newDestination
	destInfo, err = os.Stat(destination)

	if err == nil {
		if *minusOverwrite == false {
			fmt.Println("Destination file already exists!")
			os.Exit(1)
		}
	}
	// rename it
	err = os.Rename(source, destination)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
