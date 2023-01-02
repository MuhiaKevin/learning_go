package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func random(max int) int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	return rand.Intn(max)
}

func main() {
	fmt.Println("WELCOME TO THE GUESSING GAME\n")
	
	answer := random(100)
	var inputnum string

	// while loop in golang looks like this
	// For is Go's while loop

	for true {
		fmt.Print("Guess a number : ")
		fmt.Scanln(&inputnum)
		intInput, _ := strconv.Atoi(inputnum)

		if intInput == answer {
			fmt.Printf("You guessed right the answer is %v !\n" , answer)
			break
		} else if intInput < answer {
			fmt.Println("Sorry, your guess is low")

		} else if intInput > answer {
			fmt.Println("Sorry, your guess is too high")
		}
	}

}
