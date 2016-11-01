package bstree

type BSTree struct {
	rootNode *Node
}

func (bsTree *BSTree) init(value int) bool {
	bsTree.rootNode = &Node{value: value}
	return true
}

type Node struct {
	leftChild  *Node
	rightChild *Node
	value      int
}
