package bstree

import "testing"

func TestInit(t *testing.T) {
	rootVal := 100

	bsTree := &BSTree{}
	bsTree.init(rootVal)

	if bsTree.rootNode == nil || bsTree.rootNode.value != 100 {
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	tree := &BSTree{}
	treeValues := []int{4, 7, 2, 8, 1, 9, 5}

	for _, value := range treeValues {
		tree.insert(nil, value)
		t.Logf("%+v", *tree)
	}

	if tree.rootNode == nil {
		t.Fail()
	}

}
