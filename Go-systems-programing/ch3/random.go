package main

import(
	"fmt"
	"math/rand"
)
func random(min, max int) int{
	return rand.Intn(max-min) + min 
}

func main(){
	fmt.Println(random(12, 100))
}