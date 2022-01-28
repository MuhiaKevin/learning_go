package main

import (
	"fmt"
	calc "github.com/MuhiaKevin/nummanip/calc"
	calcNew "github.com/MuhiaKevin/nummanip/v3/calc"
)

	
func main() {
	result := calc.Add(12,34,4)
	fmt.Println("result =>", result )

	err, newResult := calcNew.Add()

	if err != nil {
		fmt.Println("Error => ", err)
	} else {
		fmt.Println(newResult)
	}


}