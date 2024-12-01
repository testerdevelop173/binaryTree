package main

import "fmt"

type nodeTree struct {
	Data  int
	Left  *nodeTree
	Right *nodeTree
}

func insert(root *nodeTree, value int) *nodeTree {
	if root == nil {
		return &nodeTree{Data: value}
	}
	if value < root.Data {
		root.Left = insert(root.Left, value)
	} else if value > root.Data {
		root.Right = insert(root.Right, value)
	}
	return root
}

func height(root *nodeTree) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}
func main() {
	root := &nodeTree{Data: 2}
	insert(root, 1)
	insert(root, 2)
	insert(root, 3)
	insert(root, 4)
	insert(root, 5)
	insert(root, 6)
	insert(root, 7)
	fmt.Println("Height of the binary tree:", height(root))
}
