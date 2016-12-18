package bstree

import (
	"fmt"
)

const (
	INSERT1 = iota
	INSERT2

	TRAV_PRE = iota
	TRAV_IN
	TRAV_POST
)

type BSTree struct {
	rootNode *Node

	insertMode int
	insert1    func(*Node, int) bool
	insert2    func(*Node, int) bool

	traverseMode  int
	traverserPre  func(*Node, []*Node)
	traverserIn   func(*Node, []*Node)
	traverserPost func(*Node, []*Node)
}

// Init a tree, set insertMode and traverseMode
func (bsTree *BSTree) Init(insertMode, traverseMode int) bool {
	bsTree.insertMode = insertMode
	bsTree.insert1 = firstInserter
	bsTree.insert2 = secondInserter

	bsTree.traverseMode = traverseMode
	bsTree.traverserPre = pre
	bsTree.traverserIn = in
	bsTree.traverserPost = post
	return true
}

func (bsTree *BSTree) Insert(value int) (ret bool) {
	fmt.Printf("Inserting rootNode:%+v , value:%d\n", bsTree.rootNode, value)
	if bsTree.rootNode == nil {
		bsTree.rootNode = &Node{value: value}
		fmt.Printf("set rootNode %d\n", bsTree.rootNode.value)
		ret = true
	} else if bsTree.insertMode == INSERT1 {
		//ret = bsTree.insert1(bsTree.rootNode, value)
		ret = firstInserter(bsTree.rootNode, value)
	} else if bsTree.insertMode == INSERT2 {
		ret = secondInserter(bsTree.rootNode, value)
	}
	return
}

func (bsTree *BSTree) Traverse() (ret []*Node) {
	switch bsTree.traverseMode {
	case TRAV_PRE:
		pre(bsTree.rootNode, ret)
	case TRAV_IN:
		in(bsTree.rootNode, ret)
	case TRAV_POST:
		post(bsTree.rootNode, ret)
	}
	return
}

type Node struct {
	leftChild  *Node
	rightChild *Node
	value      int
}
