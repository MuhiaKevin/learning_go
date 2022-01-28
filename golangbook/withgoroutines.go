package main

import "fmt"
//  hyperfine  './withgoroutines' './withoutgoroutines'
 
func f(n int) {
	for i := 0; i < 1000000; i++ {
		fmt.Println("fn function", n, ":", i)
	}
}

func f2(n int) {
	for i := 0; i < 1000000; i++ {
		fmt.Println("fn2 function", n, ":", i)
	}
}

func f3(n int) {
	for i := 0; i < 1000000; i++ {
		fmt.Println("fn3 function", n, ":", i)
	}
}

func main() {
	go f(0)
	go f2(0)
	go f3(0)

}
