package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}
	filename := os.Args[1]

	// myNameByte := []byte("My name is Muhia Kahiga")
	// ioutil.WriteFile(filename, myNameByte, 0644)

	fille, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
    
    fmt.Println(*fille)
	//defer fille.Close()
	//anotherByteSlice := make([]byte, 100000000)
	// fmt.Println(anotherByteSlice)

	//n, err := fille.Read(anotherByteSlice)
    //fmt.Printf("Read %d bytes: %s \n", n, anotherByteSlice)
}
