package bstree

import (
	"log"
	"testing"
)

func TestInit(t *testing.T) {
	bsTree := &BSTree{}
	bsTree.Init(0, 0)

	if bsTree.insert1 == nil || bsTree.insert2 == nil || bsTree.traverserIn == nil ||
		bsTree.traverserPost == nil || bsTree.traverserPre == nil {
		t.Fail();
	}
}

func TestInsert(t *testing.T) {
	tree := &BSTree{}
	treeValues := []int{4, 7, 2, 8, 1, 9, 5}

	log.Println(tree)

	t.Logf("test")
	for _, value := range treeValues {
		tree.Insert(value)
		t.Logf("%+v", *tree.rootNode)
	}

	if tree.rootNode == nil {
		t.Fail()
	}
}

func TestTraverPre(t *testing.T) {
	tree := &BSTree{}
	tree.Init(0, 0) // insertMode: insert1; traverMode: Pre.
	treeValues := []int{4, 7, 2, 8, 1, 9, 5}

	for _, value := range treeValues {
		tree.Insert(value)
	}

	result := tree.Traverse()
	t.Log(result)
}
