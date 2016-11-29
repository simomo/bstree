package bstree

type BSTree struct {
	rootNode *Node
	insert   func(*Node, int) bool
}

func (bsTree *BSTree) Init(value int) bool {
	bsTree.rootNode = &Node{value: value}
	bsTree.insert = secondInserter
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
	return bsTree.insert(bsTree.rootNode, value)
}

type Node struct {
	leftChild  *Node
	rightChild *Node
	value      int
}
