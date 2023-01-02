package main

// TODO:  read on Go functions returning errors
// loading needed packages
import (
	"fmt"
	"os"
	"strconv"
)

// this is where all code is executed
func main() {
	argurements := os.Args
	sum := 0

	// loop through all the commandline argurements and get their sum
	for i := 0; i < len(argurements); i++ {
		temp, _ := strconv.Atoi(argurements[i]) // interger to string conversion
		sum = sum + temp
	}
	fmt.Println("The sum is : ", sum)
}
