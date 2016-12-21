package bstree

import (
	"log"
	"testing"
	"fmt"
)

func TestInit(t *testing.T) {
	bsTree := &BSTree{}
	bsTree.Init(INSERT2, TRAV_IN)

	if bsTree.insertMode != INSERT2 || bsTree.traverseMode != TRAV_IN {
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
	tree.Init(INSERT1, TRAV_IN) // insertMode: insert1; traverMode: Pre.
	treeValues := []int{4, 7, 2, 8, 1, 9, 5}

	for _, value := range treeValues {
		tree.Insert(value)
	}

	result := tree.Traverse()
	for _, node := range result {
		fmt.Printf("Node: %+v\n", node)
	}
	t.Log(result)
}
