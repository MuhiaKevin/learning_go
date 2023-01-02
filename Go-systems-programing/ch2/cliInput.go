package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	iflag := false

	for i := 0; i < len(arguments); i++ {
		if strings.Compare(arguments[i], "-i") == 0 {
			iflag = true
			break
		}
	}

	if iflag {
		fmt.Println("Got the -i parameter !")
		fmt.Print("y/n: ")
		var answer string
		fmt.Scanln(&answer)
		fmt.Println("You entered : ", answer)
	} else {
		fmt.Println("The -i parameter is not set")
	}

}
