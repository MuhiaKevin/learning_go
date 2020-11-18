package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
	}
	if v == t.Value {
		fmt.Println("Node already exists", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	return addNode(t.Next, v)
}

func traverse(t *Node) {
	// if struct is empty say its empty
	if t == nil {
		fmt.Println("-> Empty list")
		return
	}

	// loop through t
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		// update t with the next element in the linkedlist
		t = t.Next
	}
	fmt.Println()
}

var root = new(Node)

func main() {
	fmt.Println(root)
	// root = nil
	traverse(root)
	addNode(root, 1)
	// addNode(root, 1)
	traverse(root)
	addNode(root, 10)
	addNode(root, 5)
	// addNode(root, 0)
	addNode(root, 0)
	traverse(root)
	addNode(root, 100)
	traverse(root)

}
