package bstree

import "fmt"

func pre(root *Node, result []*Node) []*Node {
	result = append(result, root)
	fmt.Printf("Appended node:%+v into result:%+v\n", root, result)
	if root.leftChild != nil {
		fmt.Printf("pre traverse left tree: %+v\n", root.leftChild)
		result = pre(root.leftChild, result)
	}
	if root.rightChild != nil {
		fmt.Printf("pre traverse right tree: %+v\n", root.rightChild)
		result = pre(root.rightChild, result)
	}

	return result
}

func in(root *Node, result []*Node) []*Node {
	if root.leftChild != nil {
		fmt.Printf("pre traverse left tree: %+v\n", root.leftChild)
		result = in(root.leftChild, result)
	}
	result = append(result, root)
	fmt.Printf("Appended node:%+v into result:%+v\n", root, result)
	if root.rightChild != nil {
		fmt.Printf("pre traverse right tree: %+v\n", root.rightChild)
		result = in(root.rightChild, result)
	}

	return result
}

func post(root *Node, result []*Node) []*Node {
	if root.leftChild != nil {
		fmt.Printf("pre traverse left tree: %+v\n", root.leftChild)
		result = in(root.leftChild, result)
	}
	if root.rightChild != nil {
		fmt.Printf("pre traverse right tree: %+v\n", root.rightChild)
		result = in(root.rightChild, result)
	}
	result = append(result, root)
	fmt.Printf("Appended node:%+v into result:%+v\n", root, result)

	return result
}
