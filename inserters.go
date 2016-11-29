package bstree

import "fmt"

func firstInserter(parent *Node, value int) bool {
	if value > parent.value {
		if parent.rightChild == nil {
			parent.rightChild = &Node{value: value}
			fmt.Printf("      %d -> %d\n", parent.value, value)
		} else {
			firstInserter(parent.rightChild, value)
		}
	} else if value < parent.value {
		if parent.leftChild == nil {
			parent.leftChild = &Node{value: value}
			fmt.Printf("%d <- %d\n", value, parent.value)
		} else {
			firstInserter(parent.leftChild, value)
		}
	} else {
		println("same value detected %s", value)
	}
	return true
}

func secondInserter(parent *Node, value int) bool {
	fmt.Printf("%+v\n", *parent)
	candidatePosition := &parent.leftChild
	arrow := "<-"
	if value > parent.value {
		candidatePosition = &parent.rightChild
		arrow = "->"
	}

	if *candidatePosition == nil {
		*candidatePosition = &Node{value: value}
		fmt.Printf("%d %s %d\n", parent.value, arrow, value)
	} else {
		secondInserter(*candidatePosition, value)
	}

	return true
}
