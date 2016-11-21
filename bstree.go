package bstree

import "fmt"

type BSTree struct {
	rootNode *Node
}

func (bsTree *BSTree) Init(value int) bool {
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

func (bsTree *BSTree) Insert(value int) bool {
	if bsTree.rootNode == nil {
		return bsTree.Init(value)
	}
	return bsTree._insert(bsTree.rootNode, value)
}

func (bsTree *BSTree) _insert(parent *Node, value int) bool {
	fmt.Printf("%+v", *parent)
	candidatePosition := parent.leftChild
	if value > parent.value {
		candidatePosition = parent.rightChild
	}

	if candidatePosition == nil {
		candidatePosition = &Node{value: value}
		println("%s\n|\n%s", parent.value, value)
	} else {
		bsTree._insert(candidatePosition, value)
	}

	return true
}

type Node struct {
	leftChild  *Node
	rightChild *Node
	value      int
}
