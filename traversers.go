package bstree

import "fmt"

func pre(root *Node, result *[]*Node) {
	result = (append(*result, root))
	fmt.Printf("Appended node:%+v into result:%+v\n", root, result)
	if root.leftChild != nil {
		fmt.Printf("pre traverse left tree: %+v\n", root.leftChild)
		pre(root.leftChild, result)
		//fmt.Printf("left tree over\n")
	}
	if root.rightChild != nil {
		fmt.Printf("pre traverse right tree: %+v\n", root.rightChild)
		pre(root.rightChild, result)
	}
}

func in(root *Node, result []*Node) {
}

func post(root *Node, result []*Node) {
}
