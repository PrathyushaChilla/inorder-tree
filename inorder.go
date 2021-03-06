package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func InorderRecursive(root *Node) {
	if root == nil {
		return
	}

	InorderRecursive(root.left)
	fmt.Printf("%d ", root.val)
	InorderRecursive(root.right)
}

func main() {
	/*
				10
			   /  \
			 20	   30
			/ \      \
		   40  50     60
		  /         /
		 70        80
	*/

	root := &Node{10, nil, nil}
	root.left = &Node{20, nil, nil}
	root.right = &Node{30, nil, nil}
	root.left.left = &Node{40, nil, nil}
	root.left.right = &Node{50, nil, nil}
	root.right.right = &Node{60, nil, nil}
	root.left.left.left = &Node{70, nil, nil}

	fmt.Println("Inorder Traversal - recursive solution : ")
	InorderRecursive(root)
}
