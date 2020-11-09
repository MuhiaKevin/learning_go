package main

import "fmt"

// https://medium.com/@riteeksrivastava/how-slices-internally-work-in-golang-a47fcb5d42ce

func main() {

	// by assigining an array to a variable a copy of the is usually made to the variable
	// array1 := [...]string{"Alice", "Bob", "Cop"}
	// array2 := a // a copy of a is assigned to array2
	// array2[0] = "Dytto"
	// fmt.Println("a is ", array1)
	// fmt.Println("b is ", array2)

	// // slices are a reference to an existing array
	// array3 := [5]int{1, 2, 3, 4, 5}
	// var slice1 []int = a[1:4] //creates a slice from a[1] to a[3]

	// slice2 := []int{1, 2, 3} //creates a array of 3 intergers and return the slice reference which is stored in a.

	//  Length is the number of elements present in the slice
	// capacity is the number of elements present in the underlying array(the array to which slice is referencing to) starting from the index from which the slice is created.

	// TODO: Modification of slices

	// readers := [4]string{"Alice", "Bob", "Charlie", "Dytto"}

	// a := readers[0:2]
	// b := readers[1:3]
	// fmt.Println("Before Modification")
	// fmt.Println("Readers", readers)
	// fmt.Println("a:", a, "b:", b)
	// // this will set the 0 element equal to "Common" change the location of the array elements "readers" which is being referenced

	// b[0] = "Common"
	// fmt.Println("After Modification")
	// fmt.Println("Readers", readers)
	// fmt.Println("a:", a, "b:", b)

	// TODO: Memory Allocation of slices

	// readers := []string{"Alice", "Bob", "Charlie"}
	// fmt.Println("Readers' array :", readers, "has old length", len(readers), "and capacity", cap(readers))
	// readers = append(readers, "Dytto")
	// fmt.Println("Readers' array :", readers, "has old length", len(readers), "and capacity", cap(readers))

	//case 1
	// a := []int{}
	// a = append(a, 1)
	// a = append(a, 2)
	// b := append(a, 3)
	// c := append(a, 4)
	// fmt.Println("Case 1")
	// fmt.Println("a: ", a, "\nb: ", b, "\nc: ", c)

	// //case 2
	// a = append(a, 3)
	// x := append(a, 4)
	// y := append(a, 5)
	// fmt.Println("Case 2")

	// fmt.Println("a: ", a, "\nx: ", x, "\ny: ", y)

	// TODO: Creating a slice with make

	// a := make([]int, 5)    // len(a) = 5 with elements as 0
	// b := make([]int, 0, 5) // len(a)=0 cap(a)=5

	// fmt.Println(a, b)

	// Create a tic-tac-toe board.
	// board := [][]string{
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// 	[]string{"_", "_", "_"},
	// }

	// // The players take turns.
	// board[0][0] = "X"
	// board[2][2] = "O"
	// board[1][2] = "X"
	// board[1][0] = "O"
	// board[0][2] = "X"

	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }

	// TODO: Appending a elements to a slice

	var s []int

	s = append(s, 0)
	s = append(s, 1)
	s = append(s, 156)
	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)

	// fmt.Println(cap(s))

	// TODO: Using Range to loop over a slices

	for index, value := range s {
		fmt.Printf("2**%d = %d\n", index, value)
	}

	fmt.Println("Ignore The index\n")

	for _, value := range s {
		fmt.Printf("%d\n", value)
	}

	fmt.Println("Ignore The value\n")

	for index, _ := range s {
		fmt.Printf("%d\n", index)
	}
}
