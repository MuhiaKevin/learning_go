package main

import "fmt"

func reverse(array)

func main() {
	arr := [9]int{234, 56, 2, 4, 767, 3, 45, 23, 8}

	i := 0
	j := len(arr) - 1

	for i < j {
	temp:
		arr[i]
		arr[i] = arr[j]
		arr[j] = temp

		i += 1
		j -= 1
	}

	fmt.Println(arr)
}
