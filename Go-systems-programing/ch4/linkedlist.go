package main

import "fmt"
// create an Node object

type Node struct {
	Value int
	Next *Node
}
var root = new(Node)

func add(t *Node, v int) int{
	if root == nil{
		t = &Node{v , nil}
	}

	if v == t.Value{
		fmt.Println("Node already exists")
		return -1
	}

	if t.Next == nil{
		t.Next = &Node{v, nil}
	}

	return add(t.Next , v)
}


func main(){
	root = nil
	add(root, 12 )
	fmt.Println(root)

}