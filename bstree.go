package bstree

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
	traverserPre  func(*BSTree) []*Node
	traverserIn   func(*BSTree) []*Node
	traverserPost func(*BSTree) []*Node
}

func (bsTree *BSTree) Init(value int) bool {
	bsTree.insert1 = firstInserter
	bsTree.insert2 = secondInserter

	bsTree.traverserPre = pre
	bsTree.traverserIn = in
	bsTree.traverserPost = post
	return true
}

func (bsTree *BSTree) Insert(value int) bool {
	if bsTree.rootNode == nil {
		bsTree.rootNode = &Node{value: value}
		return true
	}
	return bsTree.insert(bsTree.rootNode, value)
}

func (bsTree *BSTree) Traverse()

type Node struct {
	leftChild  *Node
	rightChild *Node
	value      int
}
