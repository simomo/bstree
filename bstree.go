package bstree

import "fmt"

type BSTree struct {
	rootNode *Node
}

func (bsTree *BSTree) init(value int) bool {
	bsTree.rootNode = &Node{value: value}
	return true
}

func (bsTree *BSTree) addNode(parent *Node, value int) {
	if value < parent.value {
		parent.leftChild = &Node{value: value}
	} else if value > parent.value {
		parent.rightChild = &Node{value: value}
	} else {
		println("same value detected %s", value)
	}
}

func (bsTree *BSTree) insert(parent *Node, value int) bool {
	if parent == nil {
		bsTree.init(value)
	} else {
		fmt.Printf("%+v", *parent)
		candidatePosition := parent.leftChild
		if value > parent.value {
			candidatePosition = parent.rightChild
		}

		if candidatePosition == nil {
			candidatePosition = &Node{value: value}
			println("%s\n|\n%s", parent.value, value)
		} else {
			bsTree.insert(candidatePosition, value)
		}
	}

	return true
}

type Node struct {
	leftChild  *Node
	rightChild *Node
	value      int
}