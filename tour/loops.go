package main

import "fmt"

// factorial using  a loop

func forloopFactorial(x int) int {
	result := x

	for i := 0; i < x; i++ {
		if i != 0 {
			result *= i
		}
	}

	return result
}

// while loop factorial
func whileloopFactorial(x int) int {
	result := x
	count := 1

	for count < x {
		if count != 0 {
			result *= count
		}
		count++
	}

	return result
}

// infinite for loop
func foreverForloop() {
	count := 1
	for {
		count++
		fmt.Println(count)
	}
}

func main() {

	// number := 10
	// fmt.Printf(" for loop for %v is : %d\n", number, forloopFactorial(number))
	// fmt.Printf(" while loop for %v is : %d", number, whileloopFactorial(number))
	foreverForloop()

}
