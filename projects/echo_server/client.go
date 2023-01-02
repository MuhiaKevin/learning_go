package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	location := "example.com:80"
	c, err := net.Dial("tcp", location)

	if err != nil {
		fmt.Println(err)
	}

	request := `GET / HTTP/1.0
    Host: example.com
	
	`
	c.Write([]byte(request))
	defer c.Close()
	println("Sent GET to server")
	reply := make([]byte, 1024)

	_, err = c.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}
	fmt.Println(string(reply))
}
