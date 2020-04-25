package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(x, y int, ans *int) {
	*ans = x + y
}

func subtract(x, y int, ans *int) {
	*ans = x - y
}
func divide(x, y int, ans *int) {
	*ans = x / y
}

func multiply(x, y int, ans *int) {
	*ans = x * y
}

func main() {
	args := os.Args
	var answer int

	if len(args) != 4 {
		fmt.Println("Usage: calc number operation number")
		fmt.Println("Example: calc 12 * 78 = 936")
		os.Exit(1)
	}
	firstnum, _ := strconv.Atoi(args[1])
	operation := args[2]
	secondnum, _ := strconv.Atoi(args[3])

	switch {
		case operation == "+":
			add(firstnum, secondnum, &answer)
		case operation == "-":
			subtract(firstnum, secondnum, &answer)
		case operation == "/":
			divide(firstnum, secondnum, &answer)
		case operation == "*":
			multiply(firstnum, secondnum, &answer)
	}
	

	fmt.Println(answer )

}
