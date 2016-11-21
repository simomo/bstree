package bstree

import "testing"

func TestInit(t *testing.T) {
	rootVal := 100

	bsTree := &BSTree{}
	bsTree.Init(rootVal)

	if bsTree.rootNode == nil || bsTree.rootNode.value != 100 {
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	tree := &BSTree{}
	treeValues := []int{4, 7, 2, 8, 1, 9, 5}

	t.Logf("test")
	for _, value := range treeValues {
		tree.Insert(value)
		t.Logf("%+v", *tree)
	}

	if tree.rootNode == nil {
		t.Fail()
	}

}
