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
