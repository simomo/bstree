package bstree

func pre(root *Node, result []*Node) {
	result = (append(result, root))
	if root.leftChild != nil {
		pre(root.leftChild, result)
	} else if root.rightChild != nil {
		pre(root.rightChild, result)
	}
}

func in(root *Node, result []*Node) {
}

func post(root *Node, result []*Node) {
}
